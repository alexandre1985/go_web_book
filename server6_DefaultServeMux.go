package main

import (
	"fmt"
	"log"
	"net/http"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go!")
}

func main() {

	http.HandleFunc("/welcome", messageHandler)

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
