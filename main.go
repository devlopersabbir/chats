package main

import (
	"fmt"
	"net/http"

	handler "github.com/devlopersabbir/chats/handler"
)

func main() {
	http.HandleFunc("/ws", handler.SocketHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
