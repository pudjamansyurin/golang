package poker_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	poker "example.com/server"
)

var (
	dummyGame         = &poker.SpyGame{}
	dummyBlindAlerter = &poker.SpyBlindAlerter{}
	dummyPlayerStore  = &poker.StubPlayerStore{}
	// dummyStdIn        = &bytes.Buffer{}
	dummyStdOut = &bytes.Buffer{}
)

var alert5Cases = []poker.ScheduleAlert{
	{0 * poker.TimeUnit, 100},
	{10 * poker.TimeUnit, 200},
	{20 * poker.TimeUnit, 300},
	{30 * poker.TimeUnit, 400},
	{40 * poker.TimeUnit, 500},
	{50 * poker.TimeUnit, 600},
	{60 * poker.TimeUnit, 800},
	{70 * poker.TimeUnit, 1000},
	{80 * poker.TimeUnit, 2000},
	{90 * poker.TimeUnit, 4000},
	{100 * poker.TimeUnit, 8000},
}

var alert7Cases = []poker.ScheduleAlert{
	{0 * poker.TimeUnit, 100},
	{12 * poker.TimeUnit, 200},
	{24 * poker.TimeUnit, 300},
	{36 * poker.TimeUnit, 400},
}

func TestCLI(t *testing.T) {
	t.Run("record Puja wins from input", func(t *testing.T) {
		in := userSends("5", "Puja wins")
		playerStore := &poker.StubPlayerStore{}

		game := poker.NewTexasHoldem(dummyBlindAlerter, playerStore)
		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Puja")
	})

	t.Run("record Erawan wins from input", func(t *testing.T) {
		in := userSends("5", "Erawan wins")
		playerStore := &poker.StubPlayerStore{}

		game := poker.NewTexasHoldem(dummyBlindAlerter, playerStore)
		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Erawan")
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

		assertGameStartedWith(t, game, 7)
	})

	t.Run("start game with 3 players and finish game with 'Puja' as winner", func(t *testing.T) {
		game := &poker.SpyGame{}
		stdout := &bytes.Buffer{}

		in := userSends("3", "Puja wins")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Puja")
	})

	t.Run("start game with 8 players and record 'Kusuma' as winner", func(t *testing.T) {
		game := &poker.SpyGame{}

		in := userSends("8", "Kusuma wins")
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Kusuma")
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

		game.Start(5, ioutil.Discard)
		cases := alert5Cases

		checkSchedulingCases(cases, t, blindAlerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(7, ioutil.Discard)
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
	assertPlayerWin(t, store, winner)
}

func assertPlayerWin(t *testing.T, store *poker.StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Errorf(`got %d call of RecordWin, want %d`, len(store.WinCalls), 1)
	}

	if store.WinCalls[0] != winner {
		t.Errorf(`got %q winner, want %q`, store.WinCalls[0], winner)
	}
}

func assertScheduleAlert(t *testing.T, got, want poker.ScheduleAlert) {
	t.Helper()
	if got.Amount != want.Amount {
		t.Errorf(`got amount %d, want %d`, got.Amount, want.Amount)
	}

	if got.At != want.At {
		t.Errorf(`got scheduled %v, want %v`, got.At, want.At)
	}
}

func assertGameNotStarted(t *testing.T, game *poker.SpyGame) {
	t.Helper()

	if game.StartCalled {
		t.Errorf("game should not run")
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

func checkSchedulingCases(cases []poker.ScheduleAlert, t *testing.T, alerter *poker.SpyBlindAlerter) {
	for i, want := range cases {
		testName := fmt.Sprintf(`%d scheduled for %v`, want.Amount, want.At)
		t.Run(testName, func(t *testing.T) {
			if len(alerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, alerter.Alerts)
			}

			got := alerter.Alerts[i]
			assertScheduleAlert(t, got, want)
		})
	}
}

func userSends(in ...string) *strings.Reader {
	input := strings.Join(in, "\n")
	return strings.NewReader(input)
}
