package handlers

import "go-chat-app/constants"

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

}

func handleSocketEvents(client *Client, event SocketEvent) {

}
