package pieces

import (
	"testing"
)

func TestKnightPossibleMoves(t *testing.T) {
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
				{File: B, Rank: 3},
				{File: C, Rank: 2},
			},
		},
		{
			name: "Test 2",
			args: TestingArgs{
				position: Position{
					File: D,
					Rank: 5,
				},
			},
			want: []Position{
				{File: B, Rank: 4},
				{File: B, Rank: 6},
				{File: C, Rank: 3},
				{File: C, Rank: 7},
				{File: E, Rank: 3},
				{File: E, Rank: 7},
				{File: F, Rank: 4},
				{File: F, Rank: 6},
			},
		},
	}

	for _, scenario := range tests {
		t.Run(scenario.name, func(t *testing.T) {
			knight := Knight{White}
			got := knight.GetPossibleMoves(scenario.args.position)

			// Check if each element in tt.want is contained in got
			AssertEqualPositions(t, scenario, got)
		})
	}
}
