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
					File: A,
					Rank: 1,
				},
			},
			want: []Position{
				{File: A, Rank: 2},
				{File: A, Rank: 3},
				{File: A, Rank: 4},
				{File: A, Rank: 5},
				{File: A, Rank: 6},
				{File: A, Rank: 7},
				{File: A, Rank: 8},
				{File: B, Rank: 1},
				{File: C, Rank: 1},
				{File: D, Rank: 1},
				{File: E, Rank: 1},
				{File: F, Rank: 1},
				{File: G, Rank: 1},
				{File: H, Rank: 1},
				{File: B, Rank: 2}, // diagonal moves
				{File: C, Rank: 3},
				{File: D, Rank: 4},
				{File: E, Rank: 5},
				{File: F, Rank: 6},
				{File: G, Rank: 7},
				{File: H, Rank: 8},
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
