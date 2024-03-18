package pieces

import (
	"testing"
)

func TestBishopPossibleMoves(t *testing.T) {
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
				{X: 1, Y: 1}, // diagonal moves
				{X: 2, Y: 2},
				{X: 3, Y: 3},
				{X: 4, Y: 4},
				{X: 5, Y: 5},
				{X: 6, Y: 6},
				{X: 7, Y: 7},
			},
		},
		{
			name: "Test 2",
			args: TestingArgs{
				position: Position{
					X: 7,
					Y: 7,
				},
			},
			want: []Position{
				{X: 6, Y: 6}, // diagonal moves
				{X: 5, Y: 5},
				{X: 4, Y: 4},
				{X: 3, Y: 3},
				{X: 2, Y: 2},
				{X: 1, Y: 1},
				{X: 0, Y: 0},
			},
		},
	}

	for _, scenario := range tests {
		t.Run(scenario.name, func(t *testing.T) {
			bishop := Bishop{White}
			got := bishop.GetPossibleMoves(scenario.args.position)

			// Check if each element in tt.want is contained in got
			AssertEqualPositions(t, scenario, got)
		})
	}
}
