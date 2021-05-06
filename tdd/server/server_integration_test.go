package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	player := "Puja"
	wins := 3

	for i := 0; i < wins; i++ {
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	}

	t.Run(`get score`, func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, newGetScoreRequest(player))

		assertBody(t, res.Body.String(), strconv.Itoa(wins))
		assertCode(t, res.Code, http.StatusOK)
	})

	t.Run(`get league`, func(t *testing.T) {
		want := []Player{
			{player, wins},
		}

		res := httptest.NewRecorder()
		server.ServeHTTP(res, newLeagueRequest())

		got := getLeagueFromBody(t, res.Body)
		assertLeague(t, got, want)
		assertCode(t, res.Code, http.StatusOK)
	})
}
