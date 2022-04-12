package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"poc/core"
	"poc/handlers"
	_ "poc/services"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var rdb *redis.Client

var ctx = context.Background()

// Map of all the currently active clients (websocket)
var clients = make(map[*websocket.Conn]bool)

// Map for all clients (updated)
var clientsMap core.ClientsMapType = make(map[*websocket.Conn]core.ClientNode)

// Single channel to send and receive ChatMessage data
var broadcaster = make(chan string) // communicates via username

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
		fmt.Fprintf(w, "Welcome!")
	}).Methods("GET")

	// To handle new websocket connection
	h := handlers.NewSocketHandlers(clients, broadcaster, clientsMap)
	r.HandleFunc("/ws", h.NewConnectionHandler)
	go h.HandleBroadcasts()

	log.Printf("Server starting at :%v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}

func initiateRedisClient(url string) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if err := rdb.Set(ctx, "key", "value", 0).Err(); err != nil {
		panic(err)
	} else {
		log.Println("Connected to REDIS successfully!")
	}
}
