package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func MakeRequest(method, url, name string, req, res interface{}) (int, *string, error) {
	client := &http.Client{}
	//url := "http://localhost:5050/instructor-login/nikhilsaji200@gmail.com/Soja!@1122"
	if method == "GET" {
		request, err3 := http.NewRequest(http.MethodGet, url, nil)
		request.Header.Add("Content-Type", "application/json")
		if err3 != nil {
			log.Println(err3.Error())
			return 0, nil, err3
		}
		res1, err1 := client.Do(request)
		if err1 != nil {
			log.Println(err1.Error())
			return 0, nil, err1
		}
		content, err2 := io.ReadAll(res1.Body)
		if err2 != nil {
			log.Println(err2.Error())
			return 0, nil, err2
		}
		var res2 string
		json.Unmarshal(content, &res2)
		log.Println(res2)
		return 200, &res2, nil

	}

	return 200, nil, nil
}
