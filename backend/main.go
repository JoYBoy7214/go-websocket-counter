package main

import (
	"net/http"
)

func main() {

	hub := NewHub()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		//log.Println("hello")
		ServerWs(w, r, hub)
	})
	http.ListenAndServe(":8080", nil)
}
