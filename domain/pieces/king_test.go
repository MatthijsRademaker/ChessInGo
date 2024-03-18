package pieces

import (
	"testing"
)

func TestKingPossibleMoves(t *testing.T) {
	tests := [2]TestingScenario{
		{
			name: "Test 1",
			args: TestingArgs{
				position: Position{
					X: 0,
					Y: 0,
				},
			},
			want: []Position{
				{X: 1, Y: 1},
				{X: 1, Y: 0},
				{X: 0, Y: 1},
			},
		},
		{
			name: "Test 2",
			args: TestingArgs{
				position: Position{
					X: 4,
					Y: 4,
				},
			},
			want: []Position{
				{X: 4, Y: 5}, // diagonal moves
				{X: 4, Y: 3},
				{X: 5, Y: 5},
				{X: 5, Y: 4},
				{X: 5, Y: 3},
				{X: 3, Y: 5},
				{X: 3, Y: 4},
				{X: 3, Y: 3},
			},
		},
	}

	for _, scenario := range tests {
		t.Run(scenario.name, func(t *testing.T) {
			king := King{White}
			got := king.GetPossibleMoves(scenario.args.position)

			// Check if each element in tt.want is contained in got
			AssertEqualPositions(t, scenario, got)
		})
	}
}
