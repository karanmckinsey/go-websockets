package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Environment variable not loaded")
	}
	var port string = os.Getenv("PORT");
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	log.Printf("Server starting at :%v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}

