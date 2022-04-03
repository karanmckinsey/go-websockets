package handlers

import "github.com/gorilla/websocket"

type SocketEvent struct {
	EventName 		string 				`json:"eventName"`
	EventPayload 	interface{}		`json:"eventPayload"` 
}

// Client is an interface between the websocket and the hub
type Client struct {
	hub 		*Client
	wsConn 	*websocket.Conn
	send 		chan SocketEvent
	userId 	string
}

// Hub maintains the set of active clients 
type Hub struct {
	clients 		map[*Client]bool 
	register 		chan *Client
	unregister	chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients			: 	make(map[*Client]bool),
		register		: 	make(chan *Client),
		unregister		: 	make(chan *Client),
	}
}

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