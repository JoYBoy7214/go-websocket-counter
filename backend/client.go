package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type message struct {
	Operation string `json:"operation"`
}

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}

func (c *Client) Listener() {
	defer func() {
		c.Hub.DeRegister <- c
		close(c.Send)
		c.Conn.Close()
	}()
	for {
		msg := &message{}
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("error : %v", err)
			break
		}
		c.Hub.BroadCast <- []byte(msg.Operation)
	}

}
func (c *Client) Writer() {
	defer func() {
		c.Conn.Close()
	}()
	for m := range c.Send {
		msg := &message{}
		if string(m) == "+" {
			msg.Operation = "+"
		} else {
			msg.Operation = "-"
		}
		err := c.Conn.WriteJSON(&msg)
		if err != nil {
			log.Println(err)
			break
		}

	}

}

func ServerWs(w http.ResponseWriter, r *http.Request, Hub *Hub) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print(err)
		return
	}
	newClient := &Client{
		Hub:  Hub,
		Conn: conn,
		Send: make(chan []byte, 256),
	}
	//log.Println("hello")
	newClient.Hub.Register <- newClient

	go newClient.Listener()
	go newClient.Writer()

}
