package main

import (
	"fmt"
	"net/http"
)

// SendMessageToSlack sends a message to slack
func SendMessageToSlack(message string) {
	fmt.Println("Sending message to slack...")

	http.Post(url, "")

	return nil
}
