// Package clients is a package for socket client
package clients

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn   *websocket.Conn
	roomID string
	send   chan []byte
}

type Room struct {
	clients map[*Client]bool
	mu      sync.Mutex
}

var Rooms = make(map[string]*Room)
var RoomsMu = sync.Mutex{}
