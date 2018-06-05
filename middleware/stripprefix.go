package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.ListenAndServe(":8080", nil)
}
