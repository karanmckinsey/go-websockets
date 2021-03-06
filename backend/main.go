package main

import (
	"fmt"
	"go-chat-app/handlers"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	// We'll need to check the origin of our connection this will allow us to make requests from our React
  	// development server to here. For now, we'll do no checking and just allow any connection
  	CheckOrigin: func(r *http.Request) bool { return true },
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	log.Println("Converting HTTP to WS")
	// upgrade the http connection to websocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	// continously listen for new incoming messages
	listener(ws)
}

func setupRoutes() {
	hub := handlers.NewHub()
	go hub.Run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome!")
	})
	http.HandleFunc("/ws", serveWS)
	// http.HandleFunc("/ws/{userId}", handlers.ConnectSocket)
	
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8000", nil)
}
