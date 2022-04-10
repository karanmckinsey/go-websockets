package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"poc/core"
	"poc/handlers"
	"poc/services"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var rdb *redis.Client

var ctx = context.Background()

// To "upgrade" incoming http requests to websocket reqs
var Upgrader 	= websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return false
	},
}
// Map of all the currently active clients (websocket)
var Clients 	= make(map[*websocket.Conn]bool)
// Single channel to send and receive ChatMessage data 
var Broadcaster = make(chan core.ChatMessage)


func main() {
	// environment variables config 
	if err := godotenv.Load(); err != nil {
		log.Fatal("Environment variable not loaded")
	}
	var port string = os.Getenv("PORT")
	var redisUrl string = os.Getenv("REDIS_URL")

	// redis config 
	initiateRedisClient(redisUrl)
	r := mux.NewRouter()

	// Routes 
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	// To handle new websocket connection 
	http.HandleFunc("/websocket", handlers.HandleConnections)
	// To handle messages through a subroutine 
	go services.HandleMessages()

	log.Printf("Server starting at :%v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}


func initiateRedisClient(url string) {
	rdb = redis.NewClient(&redis.Options{
        Addr:    url,
        Password: "", // no password set
        DB:       0,  // use default DB
    })
	if err := rdb.Set(ctx,"key","value",0).Err(); err != nil {
		panic(err)
	} else {
		log.Println("Connected to REDIS successfully!")
	}
}

