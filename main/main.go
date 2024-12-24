package main

import (
	"gaming/players"
	"log"
	"net/http"
)

// The Handler interface is what we need to implement in order to make a server.
// Typically, we do that by creating a struct and make it implement the interface
// by implementing its own ServeHTTP method.
// The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers.
// If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.
// type HandlerFunc has already implemented the ServeHTTP method.
// By type casting our PlayerServer function with it, we have now implemented the required Handler
func main() {
	handler := http.HandlerFunc(players.PlayerServer)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
