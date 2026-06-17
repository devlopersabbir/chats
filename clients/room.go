package clients

func GetRooom(roomID string) *Room {
	RoomsMu.Lock()

	defer RoomsMu.Unlock()

	room, exits := Rooms[roomID]
	if !exits {
		room = &Room{
			clients: make(map[*Client]bool),
		}
	}
	return room
}

func (r *Room) AddClient(c *Client) {
	r.mu.Lock()

	defer r.mu.Unlock()

	r.clients[c] = true
}

func (r *Room) RemoveClient(c *Client) {
	r.mu.Lock()

	defer r.mu.Unlock()

	delete(r.clients, c)
}
