package poker_test

import (
	"fmt"
	"strings"
	poker "tdd/server"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		player string
	}{
		{"Puja"},
		{"Kusuma"},
	}
	for _, tC := range testCases {
		testName := fmt.Sprintf(`recored %s win from input`, tC.player)
		t.Run(testName, func(t *testing.T) {
			input := fmt.Sprintf(`%s wins`, tC.player)
			in := strings.NewReader(input)
			playerStore := &poker.StubPlayerStore{}
			cli := poker.NewCLI(playerStore, in)
			cli.PlayPoker()

			poker.AssertPlayerWin(t, playerStore, tC.player)
		})
	}
}
func TestCLI(t *testing.T) {

}
