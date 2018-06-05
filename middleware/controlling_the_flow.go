package main

import (
	"log"
	"fmt"
	"net/http"
)

func middlewareFirst(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware First - Before Handler")
		next.ServeHTTP(w, r)
		log.Println("Middleware First - After Handler")
	})
}

func middlewareSecond(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware Second - Before Handler")
		if r.URL.Path == "/message" {
			if r.URL.Query().Get("password") == "pass123" {
				log.Println("You have admin access")
				next.ServeHTTP(w, r)
			} else {
				log.Println("Failed to authorize to the system")
				return
			}
		} else {
			next.ServeHTTP(w, r)
		}
		log.Println("Middleware Second - After Handler")
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing index handler")
	fmt.Fprintf(w, "Welcome!")
}

func message(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing about handler")
	fmt.Fprintf(w, "Go Middleware")
}

func iconHandler(w http.ResponseWriter, r *http.Request) {
	
}

func main() {
	http.HandleFunc("/favicon.ico", iconHandler)
	indexHandler := http.HandlerFunc(index)
	messageHandler := http.HandlerFunc(message)

	http.Handle("/", middlewareFirst(middlewareSecond(indexHandler)))
	http.Handle("/message", middlewareFirst(middlewareSecond(messageHandler)))

	server := &http.Server{
		Addr: ":8080",
	}

	log.Println("Listening...")
	server.ListenAndServe()
}
