package poker_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	poker "example.com/server"
	"github.com/gorilla/websocket"
)

func TestGETPlayers(t *testing.T) {
	store := &poker.StubPlayerStore{
		Scores: map[string]int{
			"Puja":   20,
			"Kusuma": 10,
		},
	}
	server := mustMakePlayerServer(t, store, dummyGame)

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

			assertStatus(t, res, http.StatusOK)
			assertBody(t, res.Body.String(), tC.score)
		})
	}

	t.Run("returns NotFound on missing player", func(t *testing.T) {
		req := newGetScoreRequest("Apollio")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatus(t, res, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	t.Run("it returns Accepted on POST", func(t *testing.T) {
		store := &poker.StubPlayerStore{}
		server := mustMakePlayerServer(t, store, dummyGame)

		player := "Kusuma"
		req := newPostWinRequest(player)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatus(t, res, http.StatusAccepted)
		poker.AssertPlayerWin(t, store, player)
	})
}

func TestLeague(t *testing.T) {
	t.Run("it return Ok on /league", func(t *testing.T) {
		server := mustMakePlayerServer(t, dummyPlayerStore, dummyGame)

		req := newLeagueRequest()
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatus(t, res, http.StatusOK)
		getLeagueFromResponse(t, res.Body)
	})

	t.Run(`it returns the league table as JSON`, func(t *testing.T) {
		store := &poker.StubPlayerStore{
			League: poker.League{
				{"Puja", 30},
				{"Kusuma", 27},
				{"Erawan", 55},
			},
		}
		server := mustMakePlayerServer(t, store, dummyGame)

		req := newLeagueRequest()
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatus(t, res, http.StatusOK)
		assertContentType(t, res, poker.JsonContentType)

		got := getLeagueFromResponse(t, res.Body)
		assertLeague(t, got, store.League)
	})
}

func TestGame(t *testing.T) {
	t.Run("GET /game returns OK", func(t *testing.T) {
		server := mustMakePlayerServer(t, &poker.StubPlayerStore{}, dummyGame)

		req := newGameRequest()
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatus(t, res, http.StatusOK)
	})

	t.Run("start a game with 3 players and declare 'Ruth' the winner", func(t *testing.T) {
		game := &poker.SpyGame{}
		winner := "Ruth"

		server := httptest.NewServer(mustMakePlayerServer(t, dummyPlayerStore, game))
		ws := mustDialWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")
		defer server.Close()
		defer ws.Close()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		time.Sleep(10 * time.Millisecond)
		poker.AssertGameStartedWith(t, game, 3)
		poker.AssertFinishCalledWith(t, game, winner)
	})
}

func assertContentType(t *testing.T, res *httptest.ResponseRecorder, want string) {
	t.Helper()

	if res.Result().Header.Get("content-type") != want {
		t.Errorf(`response content-type is not json, got %v`, res.Result().Header)
	}
}

func assertStatus(t *testing.T, res *httptest.ResponseRecorder, want int) {
	t.Helper()
	if res.Code != want {
		t.Errorf(`got %d, want %d`, res.Code, want)
	}
}

func assertBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf(`got %q, want %q`, got, want)
	}
}

func assertLeague(t *testing.T, got, want poker.League) {
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

func newGameRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return req
}

func getLeagueFromResponse(t *testing.T, body io.Reader) poker.League {
	t.Helper()

	league, _ := poker.NewLeague(body)
	return league
}

func mustMakePlayerServer(t *testing.T, store poker.PlayerStore, game poker.Game) *poker.PlayerServer {
	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		t.Fatal("error creating player server", err)
	}
	return server
}

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("can't open ws connection on %s %v", url, err)
	}
	return ws
}

func writeWSMessage(t *testing.T, ws *websocket.Conn, msg string) {
	t.Helper()
	if err := ws.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		t.Fatalf("can't send msg via ws connection %v", err)
	}
}
