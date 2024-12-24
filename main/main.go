package main

import (
	"gaming/players"
	"log"
	"net/http"
)

// PlayerServer implements Handler. Has a func ServeHTTP.
func main() {
	handler := &players.PlayerServer{}
	log.Fatal(http.ListenAndServe(":8080", handler))
}
