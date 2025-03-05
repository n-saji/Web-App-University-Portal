package utils

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var (
	clients   = make(map[*websocket.Conn]bool)
	clientsMu sync.Mutex // Mutex to protect the clients map
	broadcast = make(chan string)
)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	// Protect the clients map with a mutex
	clientsMu.Lock()
	clients[conn] = true
	clientsMu.Unlock()
	fmt.Println("New client connected")

	// Listen for messages from the client (optional)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
	}

	// Remove the client when it disconnects
	clientsMu.Lock()
	delete(clients, conn)
	clientsMu.Unlock()
	fmt.Println("Client disconnected")
}

func InitiateWebSockets() {
	http.HandleFunc("/ws", HandleConnections)

	for {
		msg := <-broadcast
		for client := range clients {
			if err := client.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				fmt.Println("Error writing message:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func SendMessage(msg string) {
	broadcast <- msg
}
