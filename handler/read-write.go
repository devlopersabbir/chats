package handler

import "github.com/devlopersabbir/chats/clients"

func ReadLoop(c *clients.Client, r *clients.Room) {
	defer func() {
		room.RemoveClient(c)
		c.conn.Close()
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		r.Broadcast(dc, msg)

	}
}
