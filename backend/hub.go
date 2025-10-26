package main

type Hub struct {
	Clients    map[*Client]bool
	Register   chan *Client
	DeRegister chan *Client
	BroadCast  chan []byte
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.Register:
			h.Clients[c] = true
		case c := <-h.DeRegister:
			h.Clients[c] = false
			delete(h.Clients, c)
		case message := <-h.BroadCast:
			for c := range h.Clients {
				c.Send <- message
			}
		}

	}

}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		DeRegister: make(chan *Client),
		BroadCast:  make(chan []byte),
	}
}
