package bot

import (
	"chess/domain"
	"testing"
)

type TestingScenario struct {
	name string
	args *domain.Board
	want []*domain.Board
}

func AssertEqualBoard(t *testing.T, scenario TestingScenario, got *domain.Board) {

	valid := make([]bool, len(scenario.want))
	for _, wantedBoard := range scenario.want {
		wantedBoard := wantedBoard.GetState()
		gotBoard := got.GetState()

		for y := 0; y < 8; y++ {
			for x := 0; x < 8; x++ {
				if wantedBoard[y][x] != gotBoard[y][x] {
					valid = append(valid, false)
					return
				}
			}
		}
		valid = append(valid, true)
	}

	for _, v := range valid {
		if v {
			return
		}
	}
	t.Errorf("Test %s failed.", scenario.name)
}
