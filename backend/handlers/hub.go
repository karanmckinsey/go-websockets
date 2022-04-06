package handlers

import (
	"fmt"
	"go-chat-app/constants"

	"github.com/gorilla/websocket"
)

func CreateNewSocketUser (hub *Hub, wsConn *websocket.Conn, userId string) {
	// creating a new client instance for the user 
	client := &Client{
		hub: hub,	
		wsConn: wsConn,
		send: make(chan SocketEvent), // creating a channel for this user 
		userId: userId,
	}
	fmt.Println(client)
	// invoking client methods as subroutines 

}

func NewHub() *Hub {
	return &Hub{
		clients			: 	make(map[*Client]bool),
		register		: 	make(chan *Client),
		unregister		: 	make(chan *Client),
	}
}

// Starts listening for register and unregister channels 
func (hub *Hub) Run() {
	for {
		select {
		case client := <- hub.register:
			UserRegisterEventHandler(hub, client)
		case client := <- hub.unregister:
			UserUnregesterEventHandler(hub, client)
		}
	}
}

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
