package main

import (
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type message struct {
	Operation string `json:"operation"`
}

func main() {
	const numClient = 50
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	var wg sync.WaitGroup
	wg.Add(numClient)
	for i := 0; i < numClient; i++ {
		go func(id int) {
			defer wg.Done()
			c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				log.Printf("connnection failed %d", id)
				return
			}
			defer c.Close()

			go func() {
				msg := message{}
				err := c.ReadJSON(&msg)
				if err != nil {
					log.Printf("error in reading id:%d  %v", id, err)
					return
				}
				log.Printf("message Recived from id:%d message: %v", id, msg)
			}()
			for j := 0; j < 5; j++ {
				time.Sleep(500 * time.Millisecond)
				msg := message{
					Operation: "+",
				}
				err := c.WriteJSON(&msg)
				if err != nil {
					log.Printf("Client %d write error: %v", id, err)
					return
				}
			}

		}(i)
		time.Sleep(50 * time.Millisecond)
	}

	wg.Wait()
	log.Println("✅ All clients finished testing")
}
