package players

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	getPlayerScore() int
}

type PlayerServer struct {
	winCount int
	name     string
}

func NewPlayer(name string, winCount int) *PlayerServer {
	return &PlayerServer{
		winCount: winCount,
		name:     name,
	}
}

func (p *PlayerServer) getPlayerScore(name string) int {
	return 20 // p.winCount
}

func (p *PlayerServer) ServeHTTP(response http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(response, p.getPlayerScore(playerName))
}
