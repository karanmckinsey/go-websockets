package handlers

import (
	"log"
	"net/http"
	"poc/core"

	"github.com/gorilla/websocket"
)

type SocketHandlers struct {
	clients     map[*websocket.Conn]bool // maps are reference type by default
	broadcaster chan core.Payload        // channels are reference type by default
	clientsMap  core.ClientsMapType
}

// To "upgrade" incoming http requests to websocket reqs
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func NewSocketHandlers(
	clients map[*websocket.Conn]bool,
	broadcaster chan core.Payload,
	clientsMap core.ClientsMapType,
) *SocketHandlers {
	return &SocketHandlers{clients, broadcaster, clientsMap}
}

/*
	A new connection will be:
	- assigned a new socket id
	- a map will be created with the socket vs {username, active}
	- same will be saved in the redis database (TODO)
	- every other users will be sent with updated list of active users
*/

func (s *SocketHandlers) NewConnectionHandler(w http.ResponseWriter, r *http.Request) {
	// upgrade the http request to a ws request
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	for { // infinite loop
		var payload core.Payload
		if err := ws.ReadJSON(&payload); err != nil {
			log.Println("Cannot read socket conection payload")
			log.Println(err)
			// Do return from function on error. If the function does not return,
			// then the function will spin in a tight loop printing errors.
			return
		}
		if payload.Type == "broadcast" {
			// assign new user to the map only if the payload type is broadcast (means new user)
			s.clientsMap[payload.Username] = core.ClientNodeType{
				Active:    true,
				WebSocket: ws,
			}
		}
		s.broadcaster <- payload
	}
}

// To check if the error is due to the client going away
func safeError(err error) bool {
	return websocket.IsCloseError(err, websocket.CloseGoingAway)
}

// This will be continously running in an infinite loop as a go routine from main.go
func (s *SocketHandlers) HandleBroadcasts() {
	for {
		// pick up the username for the new joined user
		payload := <-s.broadcaster
		s.messageClients(payload)
	}
}

// helper method
func (s *SocketHandlers) messageClients(payload core.Payload) {
	if payload.Username != "" {
		log.Printf("New user = %v has come, broadcasting!", payload.Username)
	}
	type node struct {
		Username string `json:"username"`
		Active   bool   `json:"active"`
	}

	switch payload.Type {
	case "direct":
		// find the user from the map with the username
		u, ok := s.clientsMap[payload.Username]
		res := core.ClientResponse{
			Message: payload.Message,
			Type:    "direct",
		}
		if ok {
			// send the message to that user
			log.Println("Sending direct message to the user", payload.Username)
			if err := u.WebSocket.WriteJSON(res); err != nil {
				log.Println("Error will writing direct message", err)
				return
			}
		}

	case "broadcast":
		// list of active users (this list will be sent to all ws connections)
		var activeUsers []string
		for uname, _ := range s.clientsMap {
			activeUsers = append(activeUsers, uname)
		}
		res := core.ClientResponse{
			Type:  "broadcast",
			Users: activeUsers,
		}
		// sending each user in the map with the latest list of active users
		for _, each := range s.clientsMap {
			if err := each.WebSocket.WriteJSON(res); err != nil {
				log.Println("Error", err)
				return
			}
		}
	}
}
