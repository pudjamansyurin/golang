package poker_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	poker "example.com/server"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	db, cleanDb := createTempFile(t, `[]`)
	defer cleanDb()

	store, err := poker.NewFileSystemPlayerStore(db)
	assertNoError(t, err)

	server := mustMakePlayerServer(t, store, dummyGame)
	player := "Puja"
	wins := 3

	for i := 0; i < wins; i++ {
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	}

	t.Run(`get score`, func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, newGetScoreRequest(player))

		assertBody(t, res.Body.String(), strconv.Itoa(wins))
		assertStatus(t, res, http.StatusOK)
	})

	t.Run(`get league`, func(t *testing.T) {
		want := poker.League{
			{player, wins},
		}

		res := httptest.NewRecorder()
		server.ServeHTTP(res, newLeagueRequest())

		got := getLeagueFromResponse(t, res.Body)
		assertLeague(t, got, want)
		assertStatus(t, res, http.StatusOK)
	})
}
