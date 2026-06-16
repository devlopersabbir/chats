// Package handler is a package for handling websocket connections
package handler

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func HandleConnection(conn *websocket.Conn) {
	for {
		// Read message from browser client
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		fmt.Println("Received message:", string(msg))
		response := "Server received: " + string(msg)

		err = conn.WriteMessage(messageType, []byte(response))
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}
