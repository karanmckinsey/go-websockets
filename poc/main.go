package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var rdb *redis.Client

var ctx = context.Background()

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
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	log.Printf("Server starting at :%v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}

