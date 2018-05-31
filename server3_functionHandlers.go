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

	mh := http.HandlerFunc(messageHandler)
	mux.Handle("/go", mh)

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
