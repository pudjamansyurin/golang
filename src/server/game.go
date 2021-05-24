package poker

import (
	"io"
	"time"
)

type Game interface {
	Start(numOfPlayers int, alertDestination io.Writer)
	Finish(winner string)
}

type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}

func NewTexasHoldem(alerter BlindAlerter, store PlayerStore) *TexasHoldem {
	return &TexasHoldem{alerter: alerter, store: store}
}

func (g *TexasHoldem) Start(numOfPlayers int, alertDestination io.Writer) {
	blindInc := time.Duration(5+numOfPlayers) * time.Second //time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind, alertDestination)
		blindTime += blindInc
	}
}

func (g *TexasHoldem) Finish(winner string) {
	g.store.RecordWin(winner)
}
