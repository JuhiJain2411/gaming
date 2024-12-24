package players

import (
	"fmt"
	"net/http"
	"strings"
)

// Strategy Pattern refers to a behavioral design pattern.
// Its task is to identify similar algorithms that solve
// a specific problem. The implementation of the algorithms
// is taken out in structs and the ability to select
// algorithms is provided at runtime

// Strategy Interface
type PlayerStore interface {
	getPlayerScore(name string) int
}

// Concrete Strategy & Strategy Context
type PlayerServer struct {
	// We give the struct a reference to the interface
	// so that we can stub the interface when testing
	// ServeHTTP and mock getPlayerScore function.
	store PlayerStore
}

/*func (p *PlayerServer) getPlayerScore(name string) int {
	return 20 // p.winCount
}*/

func (p *PlayerServer) ServeHTTP(response http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(response, p.store.getPlayerScore(playerName))
}
