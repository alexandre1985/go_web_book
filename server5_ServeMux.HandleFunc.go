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
	mux := http.NewServeMux()

	mux.HandleFunc("/go", messageHandler)

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
