package pieces

import (
	"testing"
)

func TestQueenPossibleMoves(t *testing.T) {

	tests := []TestingScenario{
		{
			name: "Test 1",
			args: TestingArgs{
				position: Position{
					X: 0,
					Y: 0,
				},
			},
			want: []Position{
				{X: 0, Y: 1},
				{X: 0, Y: 2},
				{X: 0, Y: 3},
				{X: 0, Y: 4},
				{X: 0, Y: 5},
				{X: 0, Y: 6},
				{X: 0, Y: 7},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 3, Y: 0},
				{X: 4, Y: 0},
				{X: 5, Y: 0},
				{X: 6, Y: 0},
				{X: 7, Y: 0},
				{X: 1, Y: 1}, // diagonal moves
				{X: 2, Y: 2},
				{X: 3, Y: 3},
				{X: 4, Y: 4},
				{X: 5, Y: 5},
				{X: 6, Y: 6},
				{X: 7, Y: 7},
			},
		},
	}

	for _, scenario := range tests {
		t.Run(scenario.name, func(t *testing.T) {
			queen := Queen{White}
			got := queen.GetPossibleMoves(scenario.args.position)

			// Check if each element in tt.want is contained in got

			AssertEqualPositions(t, scenario, got)
		})
	}
}
