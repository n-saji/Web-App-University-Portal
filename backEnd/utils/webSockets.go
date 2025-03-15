package utils

import (
	"CollegeAdministration/config"
	"CollegeAdministration/daos"
	"CollegeAdministration/models"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var (
	clients   = make(map[*websocket.Conn]bool)
	clientsId = make(map[string]*websocket.Conn)
	clientsMu sync.Mutex // Mutex to protect the clients map
	broadcast = make(chan string)
)

func InitiateWebSockets() {

	for {
		msg := <-broadcast
		for client := range clients {
			if err := client.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				log.Println("Error writing message:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func SendMessageAsBroadCast(msg string) {
	broadcast <- msg
}

func HandleConnections(c *gin.Context) {
	w := c.Writer
	r := c.Request
	params := c.Params
	id := params.ByName("id")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	// Protect the clients map with a mutex
	clientsMu.Lock()
	clients[conn] = true
	clientsId[id] = conn
	clientsMu.Unlock()
	log.Println("New client connected")

	// Listen for messages from the client (optional)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
	}

	// Remove the client when it disconnects
	clientsMu.Lock()
	delete(clients, conn)
	clientsMu.Unlock()
	log.Println("Client disconnected")
}

func SendMessageToClient(id string, msg string) {
	clientsMu.Lock()
	conn, ok := clientsId[id]
	if ok {
		conn.WriteMessage(websocket.TextMessage, []byte(msg))
	}
	clientsMu.Unlock()
}

func SendMessageToAllClients(msg string) {
	clientsMu.Lock()
	for conn := range clients {
		conn.WriteMessage(websocket.TextMessage, []byte(msg))
	}
	clientsMu.Unlock()
}

// message to send , account type to send to, account to skip
func SendEventToAllClients(event string, account_type, skip_account string) {

	dbConn := config.DBInit()
	db := daos.New(dbConn)

	ids, err := db.GetAccountIDsByType(account_type)
	if err != nil {
		fmt.Println("error while fetching instructor ids" + err.Error())
		return
	}
	for _, id := range ids {
		if id.Id.String() == skip_account {
			continue
		}
		msg := &models.Messages{}
		msg.ID = uuid.New()
		msg.AccountID = id.Id
		msg.Messages = event
		msg.IsRead = false
		err := db.InsertIntoMessages(msg)
		if err != nil {
			fmt.Println("error while inserting message" + err.Error())
			return
		}
	}
	clientsMu.Lock()
	for acntId,conn := range clientsId {
		if acntId == skip_account {
			continue
		}
		conn.WriteMessage(websocket.TextMessage, []byte(event))
		db.UpdateMessageStatusForAccountId(uuid.MustParse(acntId))
	}
	clientsMu.Unlock()
	defer config.CloseDB(dbConn)
}
