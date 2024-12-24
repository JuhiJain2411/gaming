package players

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	if score == 0 {
		return http.StatusNotFound
	}
	return score
}

func (s *StubPlayerStore) AddPlayerScore(name string) {
	s.scores[name]++
}

func TestPlayerGet(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Jack": 20,
			"Jill": 10,
		},
	}
	playerServer := &PlayerServer{&store}
	t.Run("Get number of GetPlayerScore for a player", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Jack", nil)
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("Wanted %v, got %v", want, got)
		}
	})

	t.Run("Missing Player", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Julie", nil)
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		got := response.Body.String()
		want := fmt.Sprint(http.StatusNotFound)

		if got != want {
			t.Errorf("Wanted %v, got %v", want, got)
		}
	})
}

func TestPlayersPost(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Jack": 20,
			"Jill": 10,
		},
	}
	playerServer := &PlayerServer{&store}
	t.Run("Record a win for a player", func(t *testing.T) {
		//POST /players/{name}
		request, _ := http.NewRequest(http.MethodPost, "/players/Julie", nil)
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)
		_, ok := store.scores["Julie"]

		// assertions
		if len(store.scores) != 3 {
			t.Errorf("Wanted %v, got %v", 3, len(store.scores))
		}
		if !ok {
			t.Errorf("Player score not added")
		}
	})
}
