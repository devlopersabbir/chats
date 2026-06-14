package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

type Hub struct {
	clients map[*Client]bool
	broadcast chan []byte

	register chan *Client
	unregister chan *Client

	mu sync.Mutex
}

var hub = Hub{
	clients: make(map[*Client]bool),

	broadcast: make(chan []byte),

	register: make(chan *Client),
	unregister: make(chan *Client),
}

func (h *Hub) run() {
	for {
		select {

		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
			log.Println("Client connected")

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			h.mu.Unlock()
			log.Println("Client disconnected")

		case message := <-h.broadcast:
			log.Println("Broadcast message: ", string(message))
			h.mu.Lock()

			for client := range h.clients {
				select {

				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
			h.mu.Unlock()
		}
	}
}

// write message to the client
func (c *Client) writePump() {
	for msg := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, msg)

		if err != nil {
			log.Println("write error: ", err)
			return
		}
	}
}
// read message from client
func (c *Client) readPump() {
	defer func() {
		// unregister 
		hub.unregister <- c
		// close connection
		c.conn.Close()
	} ()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("Read error: ", err)
		}
		hub.broadcast <- message
	}
}
// websocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all origins
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Upgrade error: ", err)
		return
	}

	client := &Client{
		conn: conn,
		send: make(chan []byte, 256),
	}

	hub.register <- client

	go client.writePump()
	go client.readPump()
}

func main() {
	go hub.run()

	http.HandleFunc("/ws", wsHandler)

	log.Println("Server started at :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
