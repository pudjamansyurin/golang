package poker

import (
	"fmt"
	"io"
	"time"
)

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.Scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
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
}

func (s *SpyGame) Start(numOfPlayers int, out io.Writer) {
	s.StartCalled = true
	s.StartCalledWith = numOfPlayers
	out.Write(s.BlindAlert)
}

func (s *SpyGame) Finish(winner string) {
	s.FinishedCalled = true
	s.FinishCalledWith = winner
}
