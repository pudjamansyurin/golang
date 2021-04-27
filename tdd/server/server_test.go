package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{
			"Puja":   20,
			"Kusuma": 10,
		},
		nil,
	}
	server := &PlayerServer{store}

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

			assertStatus(t, res.Code, http.StatusOK)
			assertResponseBody(t, res.Body.String(), tC.score)
		})
	}

	t.Run("returns 404 on missing players", func(t *testing.T) {
		req := newGetScoreRequest("Apollio")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := res.Code
		want := http.StatusNotFound

		assertStatus(t, got, want)
	})
}

func TestStoreWins(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := &PlayerServer{store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Kusuma"
		req := newPostWinRequest(player)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatus(t, res.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf(`got %d call, want %d`, len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf(`got %q winner, want %q`, store.winCalls[0], player)
		}
	})
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf(`got %d, want %d`, got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf(`got %q, want %q`, got, want)
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

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}
