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
					File: A,
					Rank: 1,
				},
			},
			want: []Position{
				{File: B, Rank: 2}, // diagonal moves
				{File: C, Rank: 3},
				{File: D, Rank: 4},
				{File: E, Rank: 5},
				{File: F, Rank: 6},
				{File: G, Rank: 7},
				{File: H, Rank: 8},
			},
		},
		{
			name: "Test 2",
			args: TestingArgs{
				position: Position{
					File: H,
					Rank: 8,
				},
			},
			want: []Position{
				{File: G, Rank: 7}, // diagonal moves
				{File: F, Rank: 6},
				{File: E, Rank: 5},
				{File: D, Rank: 4},
				{File: C, Rank: 3},
				{File: B, Rank: 2},
				{File: A, Rank: 1},
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
