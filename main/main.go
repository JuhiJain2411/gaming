package main

import (
	"gaming/players"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerStore) AddPlayerScore(name string, score int) {
}

// PlayerServer implements Handler. Has a func ServeHTTP.
func main() {
	handler := &players.PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":8080", handler))
}
