package poker

import (
	"fmt"
	"testing"
	"time"
)

func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Errorf(`got %d call, want %d`, len(store.winCalls), 1)
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

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
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

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, ScheduleAlert{duration, amount})
}

type SpyGame struct {
	Run          bool
	StartedWith  int
	FinishedWith string
}

func (s *SpyGame) Start(numOfPlayers int) {
	s.StartedWith = numOfPlayers
	s.Run = true
}

func (s *SpyGame) Finish(winner string) {
	s.FinishedWith = winner
}
