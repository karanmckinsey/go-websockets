package main 

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	// We'll need to check the origin of our connection this will allow us to make requests from our React
  	// development server to here. For now, we'll do no checking and just allow any connection
  	CheckOrigin: func(r *http.Request) bool { return true },
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	// upgrade the http connection to websocket connection 
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	// continously listen for new incoming messages 
	listener(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome!")
	})
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8000", nil)
}
