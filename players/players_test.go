package players

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayers(t *testing.T) {
	t.Run("Get number of wins for a player", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "GET /players/{name}", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

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
