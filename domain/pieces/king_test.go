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
					File: A,
					Rank: 1,
				},
			},
			want: []Position{
				{File: B, Rank: 2},
				{File: B, Rank: 1},
				{File: A, Rank: 2},
			},
		},
		{
			name: "Test 2",
			args: TestingArgs{
				position: Position{
					File: D,
					Rank: 4,
				},
			},
			want: []Position{
				{File: D, Rank: 5}, // diagonal moves
				{File: D, Rank: 3},
				{File: E, Rank: 5},
				{File: E, Rank: 4},
				{File: E, Rank: 3},
				{File: C, Rank: 5},
				{File: C, Rank: 4},
				{File: C, Rank: 3},
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
