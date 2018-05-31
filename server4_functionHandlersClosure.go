package main

import (
	"fmt"
	"log"
	"net/http"
)

func mh(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/message", mh("Hello"))
	mux.Handle("/hi", mh("Hi"))

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
