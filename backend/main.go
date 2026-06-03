package main

import (
	"log"
	"net/http"
)

func main() {
	// register handlers from lobby.go and game.go (same package)
	http.HandleFunc("/create", createGameHandler)
	http.HandleFunc("/join", joinGameHandler)
	http.HandleFunc("/ws", gameWSHandler)

	log.Println("server listening on :5000")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}
