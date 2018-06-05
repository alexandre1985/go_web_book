package main

import (
"log"
"fmt"
"net/http"

"github.com/codegangsta/negroni"
)

func middlewareFirst(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Middleware First - Before Handler")
	next(w, r)
	log.Println("Middleware First - After Handler")
}

func middlewareSecond(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Middleware Second - Before Handler")
	if r.URL.Path == "/message" {
		if r.URL.Query().Get("password") == "pass123" {
			log.Println("You have admin access")
			next(w, r)
		} else {
			log.Println("Failed to authorize to the system")
			return
		}
	} else {
		next(w, r)
	}
	log.Println("Middleware Second - After Handler")
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
	mux := http.NewServeMux()
	mux.HandleFunc("/favicon.ico", iconHandler)
	mux.HandleFunc("/", index)
	mux.HandleFunc("/message", message)

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(middlewareFirst))
	n.Use(negroni.HandlerFunc(middlewareSecond))

	/*
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(middlewareFirst),
		negroni.HandlerFunc(middlewareSecond),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
	)
	*/
	
	n.UseHandler(mux)

	n.Run(":8080")
}
