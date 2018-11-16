package main

import (
	"log"
	"net/http"

	"./handlers"
)

func main() {
	server := &http.Server{
		Addr: ":8000",
	}
	http.Handle("/", handlers.StaticHandler("assets"))
	http.HandleFunc("/submit", handlers.SubmitHandler)
	http.HandleFunc("/unknown", handlers.UnknownHandler)
	log.Fatal(server.ListenAndServe())
}
