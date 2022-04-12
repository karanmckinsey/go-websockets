package core

import "github.com/gorilla/websocket"

type ChatMessage struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type ClientNode struct {
	Username string `json:"username"`
	Active   bool   `json:"active"`
}

type NewSessionPayload struct {
	Username string `json:"username"`
}

type ClientsMapType map[*websocket.Conn]ClientNode
