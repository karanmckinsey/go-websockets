package handlers

import (
	"fmt"
	"go-chat-app/constants"
)

func UserRegisterEventHandler(hub *Hub, client *Client) {
	// setting the client to true in the hub struct
	hub.clients[client] = true
	handleSocketEvents(
		client,
		SocketEvent{
			EventName:    constants.EventTypes["join"],
			EventPayload: client.userId,
		},
	)
}

func UserUnregesterEventHandler(hub *Hub, client *Client) {
	hub.clients[client] = false
}

func handleSocketEvents(client *Client, event SocketEvent) {
	switch event.EventPayload {
		case constants.EventTypes["JOIN"]:
			fmt.Printf("Client %v has joined the hub!", client.userId)
	}
}
