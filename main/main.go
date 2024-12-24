package main

import (
	"gaming/players"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{make(map[string]int)}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) AddPlayerScore(name string) {
	i.store[name]++
}

// PlayerServer implements Handler. Has a func ServeHTTP.
func main() {
	handler := &players.PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":8080", handler))
}
