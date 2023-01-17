package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func MakeRequest(method, url, name string, req, res interface{}) (int, error) {

	client := &http.Client{}
	if method == "GET" {
		request, err3 := http.NewRequest(method, url, nil)
		request.Header.Set("Internal-call", "true")
		if err3 != nil {
			log.Println(err3.Error())
			return 0, err3
		}
		response, err1 := client.Do(request)
		if err1 != nil {
			log.Println(err1.Error())
			return 0, err1
		}
		content, err2 := io.ReadAll(response.Body)
		if err2 != nil {
			log.Println(err2.Error())
			return 0, err2
		}
		json.Unmarshal(content, &res)
		return 200, nil

	}

	return 200, nil
}
