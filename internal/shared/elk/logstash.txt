package elk

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

const addr = "0.0.0.0:5044"

func LogToLogstash(message string) {
	jsonStr := []byte(`{"message":"` + message + `"}`)
	resp, err := http.Post(fmt.Sprintf("http://%s", addr), "application/json", bytes.NewBuffer(jsonStr))
	//resp, err := http.Post("http://localhost:5044", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Printf("Failed to send log to Logstash: %v", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
}
