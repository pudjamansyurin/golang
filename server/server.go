package poker

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
	tmpl *template.Template
	game Game
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

const JsonContentType = "application/json"
const htmlTemplatePath = "game.html"

func NewPlayerServer(store PlayerStore, game Game) (*PlayerServer, error) {
	p := new(PlayerServer)

	tmpl, err := template.ParseFiles("game.html")
	if err != nil {
		return nil, fmt.Errorf("fail loading template %s %v", htmlTemplatePath, err)
	}

	p.game = game
	p.tmpl = tmpl
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	router.Handle("/game", http.HandlerFunc(p.gameHandler))
	router.Handle("/ws", http.HandlerFunc(p.wsHandler))

	p.Handler = router
	return p, nil
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", JsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.proccessWin(w, player)
	}
}

func (p *PlayerServer) gameHandler(w http.ResponseWriter, r *http.Request) {
	p.tmpl.Execute(w, nil)
}

func (p *PlayerServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	ws := newPlayerServerWS(w, r)

	numberOfPlayersMsg := ws.WaitForMsg()
	numberOfPlayers, _ := strconv.Atoi(string(numberOfPlayersMsg))
	p.game.Start(numberOfPlayers, ws)

	winnerMsg := ws.WaitForMsg()
	p.game.Finish(string(winnerMsg))
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) proccessWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
