package handlers

import (
	"log"
	"net/http"
	"poc/core"

	"github.com/gorilla/websocket"
)

type SocketHandlers struct {
	upgrader *websocket.Upgrader
	clients map[*websocket.Conn]bool // maps are reference type by default 
	broadcaster chan core.ChatMessage // channels are reference type by default 
}

func NewSocketHandlers(
	upgrader *websocket.Upgrader, 
	clients map[*websocket.Conn]bool,
	broadcaster chan core.ChatMessage,
) *SocketHandlers {
	return &SocketHandlers{upgrader, clients, broadcaster}
}

/*
	Whenever the users makes a new connection:
	- They should be set up to receive messages from other clients.
	- They should be able to send their own messages.
	- They should receive a full history of the previous chat (backed by Redis).

*/

func (s *SocketHandlers) HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Create a new WS connection out of the incoming request
	ws, err := s.upgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	// Set the currect WS as true in the map
	s.clients[ws] = true 
	// Send the message payload received with the WS connection 
	for {
		var chatMessage core.ChatMessage 
		if err := ws.ReadJSON(&chatMessage); err != nil {
			// Delete the current client from the map in case of an error 
			delete(s.clients, ws)
			break 
		} else {
			s.broadcaster <- chatMessage
		}

	}
}
