package main

import (
	"fmt"
	"log"
	"net/http"
	handler "nicotheowedding/api"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Server is running at http://localhost:2300/run")
	http.HandleFunc("/run", handler.Handler)
	log.Fatal(http.ListenAndServe(":2300", nil))
}
