package poker_test

import (
	"bytes"
	"fmt"
	"strings"
	poker "tdd/server"
	"testing"
	"time"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

var alert5Cases = []poker.ScheduleAlert{
	{0 * time.Second, 100},
	{10 * time.Minute, 200},
	{20 * time.Minute, 300},
	{30 * time.Minute, 400},
	{40 * time.Minute, 500},
	{50 * time.Minute, 600},
	{60 * time.Minute, 800},
	{70 * time.Minute, 1000},
	{80 * time.Minute, 2000},
	{90 * time.Minute, 4000},
	{100 * time.Minute, 8000},
}

var alert7Cases = []poker.ScheduleAlert{
	{0 * time.Second, 100},
	{12 * time.Minute, 200},
	{24 * time.Minute, 300},
	{36 * time.Minute, 400},
}

func TestCLI(t *testing.T) {
	t.Run("record Puja wins from input", func(t *testing.T) {
		in := userSends("5", "Puja wins")
		playerStore := &poker.StubPlayerStore{}

		game := poker.NewTexasHoldem(dummyBlindAlerter, playerStore)
		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Puja")
	})

	t.Run("record Erawan wins from input", func(t *testing.T) {
		in := userSends("5", "Erawan wins")
		playerStore := &poker.StubPlayerStore{}

		game := poker.NewTexasHoldem(dummyBlindAlerter, playerStore)
		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Erawan")
	})

	t.Run(`it schedules printing of blind values`, func(t *testing.T) {
		in := userSends("5", "Puja wins")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &poker.SpyBlindAlerter{}

		game := poker.NewTexasHoldem(blindAlerter, playerStore)
		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		checkSchedulingCases(alert5Cases, t, blindAlerter)
	})

	t.Run(`it prompts the user to enter the number of players`, func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("7")
		game := &poker.SpyGame{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted start called with 7 but got %d", game.StartedWith)
		}
	})

	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &poker.SpyGame{}
		stdout := &bytes.Buffer{}

		in := userSends("3", "Chris wins")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &poker.SpyGame{}

		in := userSends("8", "Cleo wins")
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &poker.SpyGame{}

		stdout := &bytes.Buffer{}
		in := userSends("pies")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
}
func TestGame_Start(t *testing.T) {
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(5)
		cases := alert5Cases

		checkSchedulingCases(cases, t, blindAlerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(7)
		cases := alert7Cases

		checkSchedulingCases(cases, t, blindAlerter)
	})

	t.Run("it prints an error when non numeric value is inserted", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("Pies")
		game := &poker.SpyGame{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
}

func TestGame_Finish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	game := poker.NewTexasHoldem(dummyBlindAlerter, store)
	winner := "Puja"

	game.Finish(winner)
	poker.AssertPlayerWin(t, store, winner)
}

func assertGameNotStarted(t *testing.T, game *poker.SpyGame) {
	t.Helper()

	if game.Run {
		t.Errorf("game should not run")
	}
}

func checkSchedulingCases(cases []poker.ScheduleAlert, t *testing.T, alerter *poker.SpyBlindAlerter) {
	for i, want := range cases {
		testName := fmt.Sprintf(`%d scheduled for %v`, want.Amount, want.At)
		t.Run(testName, func(t *testing.T) {
			if len(alerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, alerter.Alerts)
			}

			got := alerter.Alerts[i]
			poker.AssertScheduleAlert(t, got, want)
		})
	}
}

func userSends(in ...string) *strings.Reader {
	input := strings.Join(in, "\n")
	return strings.NewReader(input)
}

func assertFinishCalledWith(t *testing.T, game *poker.SpyGame, winner string) {
	t.Helper()
	if game.FinishedWith != winner {
		t.Errorf(`it finished with %q winner, want %q`, game.FinishedWith, winner)
	}
}

func assertGameStartedWith(t *testing.T, game *poker.SpyGame, numOfPlayers int) {
	t.Helper()
	if game.StartedWith != numOfPlayers {
		t.Errorf(`it started with %d players, want %d`, game.StartedWith, numOfPlayers)
	}
}

func assertMessagesSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()

	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}
