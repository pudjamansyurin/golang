package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := &StubPlayerStore{
		scores: map[string]int{
			"Puja":   20,
			"Kusuma": 10,
		},
		winCalls: nil,
		league:   nil,
	}
	server := NewPlayerServer(store)

	testCases := []struct {
		name  string
		score string
	}{
		{"Puja", "20"},
		{"Kusuma", "10"},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			req := newGetScoreRequest(tC.name)
			res := httptest.NewRecorder()

			server.ServeHTTP(res, req)

			assertCode(t, res.Code, http.StatusOK)
			assertBody(t, res.Body.String(), tC.score)
		})
	}

	t.Run("returns NotFound on missing player", func(t *testing.T) {
		req := newGetScoreRequest("Apollio")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertCode(t, res.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := &StubPlayerStore{}
	server := NewPlayerServer(store)

	t.Run("it returns Accepted on POST", func(t *testing.T) {
		player := "Kusuma"
		req := newPostWinRequest(player)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertCode(t, res.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf(`got %d call, want %d`, len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf(`got %q winner, want %q`, store.winCalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {
	t.Run("it return Ok on /league", func(t *testing.T) {
		store := &StubPlayerStore{}
		server := NewPlayerServer(store)

		req := newLeagueRequest()
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertCode(t, res.Code, http.StatusOK)
		getLeagueFromBody(t, res.Body)
	})

	t.Run(`it returns the league table as JSON`, func(t *testing.T) {
		store := &StubPlayerStore{
			scores:   nil,
			winCalls: nil,
			league: []Player{
				{"Puja", 30},
				{"Kusuma", 27},
				{"Erawan", 55},
			},
		}
		server := NewPlayerServer(store)

		req := newLeagueRequest()
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertCode(t, res.Code, http.StatusOK)
		assertContentType(t, res, jsonContentType)

		got := getLeagueFromBody(t, res.Body)
		assertLeague(t, got, store.league)
	})
}

func assertContentType(t *testing.T, res *httptest.ResponseRecorder, want string) {
	t.Helper()

	if res.Result().Header.Get("content-type") != want {
		t.Errorf(`response content-type is not json, got %v`, res.Result().Header)
	}
}

func assertCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf(`got %d, want %d`, got, want)
	}
}

func assertBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf(`got %q, want %q`, got, want)
	}
}

func assertLeague(t *testing.T, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`got %v, want %v`, got, want)
	}
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func getLeagueFromBody(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()

	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf(`Unable to parse players %q, %v`, body, err)
	}
	return
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}
