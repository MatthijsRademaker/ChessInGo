package pieces

import (
	"testing"
)

type TestingArgs struct {
	position Position
}
type TestingScenario struct {
	name string
	args TestingArgs
	want []Position
}

func AssertEqualPositions(t *testing.T, scenario TestingScenario, got []Position) {
	for _, wantMove := range scenario.want {
		found := false
		for _, move := range got {
			if move == wantMove {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("GetPossibleMoves() = %v, does not contain %v", got, wantMove)
		}
	}

	// Check if the length of got matches the length of tt.want
	if len(got) != len(scenario.want) {
		t.Errorf("GetPossibleMoves() = %v, want %v", got, scenario.want)
	}
}
