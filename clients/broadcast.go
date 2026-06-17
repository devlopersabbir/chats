package clients

func (r *Room) Broadcast(sender *Client, msg []byte) {
	r.mu.Lock()

	defer r.mu.Unlock()

	for client := range r.clients {
		if client != sender {
			client.send <- msg
		}
	}
}
