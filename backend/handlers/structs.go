package handlers

import "github.com/gorilla/websocket"

// Socket events structure
type SocketEvent struct {
	EventName 		string 			`json:"eventName"`
	EventPayload 	interface{}		`json:"eventPayload"` 
}

// Client is an interface between the websocket and the hub
type Client struct {
	hub 		*Client
	wsConn 		*websocket.Conn
	send 		chan SocketEvent
	userId 		string
}

// Hub maintains the set of active clients 
type Hub struct {
	clients 		map[*Client]bool 
	register 		chan *Client
	unregister		chan *Client
}