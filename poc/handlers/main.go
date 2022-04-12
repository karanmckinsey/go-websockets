package handlers

import (
	"log"
	"net/http"
	"poc/core"

	"github.com/gorilla/websocket"
)

type SocketHandlers struct {
	clients     map[*websocket.Conn]bool // maps are reference type by default
	broadcaster chan string              // channels are reference type by default
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
	broadcaster chan string,
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
		var payload core.NewSessionPayload
		if err := ws.ReadJSON(&payload); err != nil {
			log.Println("Cannot read socket conection payload")
			log.Println(err)
			return
		}
		s.clientsMap[payload.Username] = core.ClientNodeType{
			Active:    true,
			WebSocket: ws,
		}
		s.broadcaster <- payload.Username
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
		username := <-s.broadcaster
		s.messageClients(username)
	}
}

// helper method
func (s *SocketHandlers) messageClients(username string) {
	if username != "" {
		log.Printf("New user = %v has come, broadcasting!", username)
	}
	type node struct {
		Username string `json:"username"`
		Active   bool   `json:"active"`
	}
	// list of active users (this list will be sent to all ws connections)
	var activeUsers []node
	for uname, each := range s.clientsMap {
		n := node{
			Username: uname,
			Active:   each.Active,
		}
		activeUsers = append(activeUsers, n)
	}
	// sending each user in the map with the latest list of active users
	for _, each := range s.clientsMap {
		if err := each.WebSocket.WriteJSON(activeUsers); err != nil {
			log.Printf("Error")
		}
	}
}

/*
	Whenever the users makes a new connection:
	- They should be set up to receive messages from other clients.
	- They should be able to send their own messages.
	- They should receive a full history of the previous chat (backed by Redis).

*/
// NOT TO BE USED ANYMORE
// func (s *SocketHandlers) HandleConnection(w http.ResponseWriter, r *http.Request) {
// 	// Create a new WS connection out of the incoming request
// 	log.Println("Upgrading the request to ws request")
// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// defer ws.Close()
// 	// Set the currect WS as true in the map
// 	log.Println("Creating a map entry for the new connection")
// 	s.clients[ws] = true
// 	// Send the message payload received with the WS connection
// 	for {
// 		var chatMessage core.ChatMessage
// 		if err := ws.ReadJSON(&chatMessage); err != nil {
// 			// Delete the current client from the map in case of an error
// 			delete(s.clients, ws)
// 			break
// 		} else {
// 			if chatMessage.Message != "" {
// 				// Broadcast the message to the broadcast channel
// 				log.Println("Broadcasting the message", chatMessage)
// 				s.broadcaster <- chatMessage
// 			}
// 		}

// 	}
// }
