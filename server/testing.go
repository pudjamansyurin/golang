package poker

import (
	"fmt"
	"io"
	"testing"
	"time"
)

type StubPlayerStore struct {
	Scores   map[string]int
	winCalls []string
	League   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.Scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

type ScheduleAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduleAlert) String() string {
	return fmt.Sprintf(`%d chips at %v`, s.Amount, s.At)
}

type SpyBlindAlerter struct {
	Alerts []ScheduleAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	s.Alerts = append(s.Alerts, ScheduleAlert{duration, amount})
}

type SpyGame struct {
	StartCalled     bool
	StartCalledWith int
	BlindAlert      []byte

	FinishedCalled   bool
	FinishCalledWith string
	// https://quii.gitbook.io/learn-go-with-tests/build-an-application/websockets#write-the-test-first-3
}

func (s *SpyGame) Start(numOfPlayers int, alertDestination io.Writer) {
	s.StartCalledWith = numOfPlayers
	s.StartCalled = true
}

func (s *SpyGame) Finish(winner string) {
	s.FinishCalledWith = winner
	s.FinishedCalled = true
}

func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Errorf(`got %d call of RecordWin, want %d`, len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf(`got %q winner, want %q`, store.winCalls[0], winner)
	}
}

func AssertScheduleAlert(t *testing.T, got, want ScheduleAlert) {
	t.Helper()
	if got.Amount != want.Amount {
		t.Errorf(`got amount %d, want %d`, got.Amount, want.Amount)
	}

	if got.At != want.At {
		t.Errorf(`got scheduled %v, want %v`, got.At, want.At)
	}
}

func AssertFinishCalledWith(t *testing.T, game *SpyGame, winner string) {
	t.Helper()
	if game.FinishCalledWith != winner {
		t.Errorf(`it finished with %q winner, want %q`, game.FinishCalledWith, winner)
	}
}

func AssertGameStartedWith(t *testing.T, game *SpyGame, numOfPlayers int) {
	t.Helper()
	if game.StartCalledWith != numOfPlayers {
		t.Errorf(`it started with %d players, want %d`, game.StartCalledWith, numOfPlayers)
	}
}
