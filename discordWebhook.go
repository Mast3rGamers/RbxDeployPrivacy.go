package main

import (
	"fmt"
	"net/http"
	"strings"
)

func sendMessage(message string) {
	if message == "" || webhookId == "" || webhookToken == "" {
		return
	}
	reqBody := strings.NewReader(`{"content": "` + message + `"}`)

	req, err := http.NewRequest("POST", "https://discord.com/api/webhooks/"+webhookId+"/"+webhookToken, reqBody)
	req.Header.Set("Content-Type", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("got error: ", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 204 {
		fmt.Printf("couldnt send message (got %d)\n", res.StatusCode)
	}

	return
}
