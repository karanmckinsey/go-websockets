package services

import (
	"log"
	"poc/core"

	"github.com/gorilla/websocket"
)

type SocketService struct {
	broadcaster chan core.ChatMessage
	clients map[*websocket.Conn]bool
}

func NewSocketService(
	broadcaster chan core.ChatMessage,
	clients map[*websocket.Conn]bool,
	) *SocketService {
	return &SocketService{broadcaster, clients}
}

// This will be continously running in an infinite loop as a go routine from main.go
func (s *SocketService) HandleMessages() {
	for {
		// Grab any message that comes to this channel 
		msg := <- s.broadcaster
		// TODO : store this message in redis 
		// Message clients 
		messageClients(msg, s.clients)
	}
}

func  messageClients(msg core.ChatMessage, clients map[*websocket.Conn]bool) {
	for client := range clients {
		if err := client.WriteJSON(msg); err != nil && !safeError(err) {
			log.Printf("Error: %v", err)
			// Close the connection for the current client 
			client.Close()
			// Delete the client from the map 
			delete(clients, client)
		}
	}
}

// To check if the error is due to the client going away 
func safeError(err error) bool {
	return websocket.IsCloseError(err, websocket.CloseGoingAway)
}