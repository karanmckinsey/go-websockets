package main 

import (
	"fmt"
	"github.com/gorilla/websocket"
)

func listener(conn *websocket.Conn) {

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return 
		}
		fmt.Println("Message Received", string(p))
	}
	
}