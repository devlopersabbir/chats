// Package handler is a package for handling websocket connections
package handler

import (
	"fmt"
	"net/http"

	"github.com/devlopersabbir/chats/clients"
	"github.com/devlopersabbir/chats/config"
)

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("room")

	if roomID == "" {
		fmt.Println("No room ID provided")
		return
	}

	conn, err := config.Upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}

	defer conn.Close()
	fmt.Println("A Client connected")

	client := &clients.Client{
		conn:   conn,
		roomID: roomID,
		send:   make(chan []byte),
	}

	room := clients.GetRooom(roomID)
	room.AddClient(client)

	go ReadLoop(client, room)
	go WriteLoop(client, room)

	HandleConnection(conn)
}
