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
	GetPlayerScore(name string) int
	AddPlayerScore(name string)
}

// Concrete Strategy & Strategy Context
type PlayerServer struct {
	// We give the struct a reference to the interface
	// so that we can stub the interface when testing
	// ServeHTTP and mock GetPlayerScore function.
	Store PlayerStore
}

/*func (p *PlayerServer) GetPlayerScore(name string) int {
	return 20 // p.winCount
}*/

func (p *PlayerServer) ServeHTTP(response http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.Path, "/players/")
	if r.Method == http.MethodPost {
		p.Store.AddPlayerScore(playerName)
		return
	} else {
		fmt.Fprint(response, p.Store.GetPlayerScore(playerName))
	}
}
