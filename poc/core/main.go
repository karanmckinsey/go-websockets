package core

import "github.com/gorilla/websocket"

type ChatMessage struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type NewSessionPayload struct {
	Username string `json:"username"`
}

// type ClientNode struct {
// 	Username string `json:"username"`
// 	Active   bool   `json:"active"`
// }
// type ClientsMapType map[*websocket.Conn]ClientNode

type ClientNodeType struct {
	WebSocket *websocket.Conn `json:"-"`
	Active    bool            `json:"active"`
}
type ClientsMapType map[string]ClientNodeType
