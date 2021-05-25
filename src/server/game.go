package poker

import (
	"io"
	"time"
)

const TimeUnit = time.Second

type Game interface {
	Start(numOfPlayers int, out io.Writer)
	Finish(winner string)
}

type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}

func NewTexasHoldem(alerter BlindAlerter, store PlayerStore) *TexasHoldem {
	return &TexasHoldem{alerter: alerter, store: store}
}

func (g *TexasHoldem) Start(numOfPlayers int, out io.Writer) {
	blindInc := time.Duration(5+numOfPlayers) * TimeUnit

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * TimeUnit
	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind, out)
		blindTime += blindInc
	}
}

func (g *TexasHoldem) Finish(winner string) {
	g.store.RecordWin(winner)
}
