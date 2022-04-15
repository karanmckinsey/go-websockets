package core

import "github.com/gorilla/websocket"

type ChatMessage struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type Payload struct {
	Type     string `json:"type,omitempty"` // broadcast | direct
	Username string `json:"username"`
	Message  string `json:"message,omitempty"`
}

type ClientNodeType struct {
	WebSocket *websocket.Conn `json:"-"`
	Active    bool            `json:"active"`
}

type ClientsMapType map[string]ClientNodeType

type ClientResponse struct {
	Type    string   `json:"type"` // broadcast | direct
	Message string   `json:"message,omitempty"`
	Users   []string `json:"users,omitempty"`
}

var x ClientResponse = ClientResponse{
	Type: "sdt",
}
