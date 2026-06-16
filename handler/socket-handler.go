// Package handler is a package for handling websocket connections
package handler

import (
	"fmt"
	"net/http"

	"github.com/devlopersabbir/chats/config"
)

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := config.Upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}

	defer conn.Close()
	fmt.Println("A Client connected")

	HandleConnection(conn)
}
