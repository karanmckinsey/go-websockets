package handlers

import (
	"log"
	"net/http"
	"poc/core"

	"github.com/gorilla/websocket"
)

// To "upgrade" incoming http requests to websocket reqs
var upgrader 	= websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

type SocketHandlers struct {
	clients map[*websocket.Conn]bool // maps are reference type by default 
	broadcaster chan core.ChatMessage // channels are reference type by default 
}

func NewSocketHandlers(
		clients map[*websocket.Conn]bool,
		broadcaster chan core.ChatMessage,
	) *SocketHandlers {
	return &SocketHandlers{clients, broadcaster}
}


/*
	Whenever the users makes a new connection:
	- They should be set up to receive messages from other clients.
	- They should be able to send their own messages.
	- They should receive a full history of the previous chat (backed by Redis).

*/

func (s *SocketHandlers) HandleConnection(w http.ResponseWriter, r *http.Request) {
	// Create a new WS connection out of the incoming request
	log.Println("Upgrading the request to ws request")
	ws, err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Fatal(err)
	}
	// defer ws.Close()
	// Set the currect WS as true in the map
	log.Println("Creating a map entry for the new connection")
	s.clients[ws] = true 
	// Send the message payload received with the WS connection 
	for {
		var chatMessage core.ChatMessage 
		if err := ws.ReadJSON(&chatMessage); err != nil {
			// Delete the current client from the map in case of an error 
			delete(s.clients, ws)
			break 
		} else {
			if chatMessage.Message != "" {
				// Broadcast the message to the broadcast channel
				log.Println("Broadcasting the message", chatMessage)
				s.broadcaster <- chatMessage
			}
		}

	}
}
