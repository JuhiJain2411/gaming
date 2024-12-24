package players

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) getPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Jack": 20,
			"Jill": 10,
		},
	}
	playerServer := &PlayerServer{&store}
	t.Run("Get number of getPlayerScore for a player", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Jack", nil)
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("Wanted %v, got %v", want, got)
		}
	})

	t.Run("Record a win for a player", func(t *testing.T) {
		//POST /players/{name}

	})
}
