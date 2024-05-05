package pieces

import (
	"testing"
)

func (p *Rook) TestPossibleMoves(t *testing.T) {
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
			},
		},
		{
			name: "Test 2",
			args: TestingArgs{
				position: Position{
					File: D,
					Rank: 3,
				},
			},
			want: []Position{
				{File: D, Rank: 1},
				{File: D, Rank: 1},
				{File: D, Rank: 2},
				{File: D, Rank: 4},
				{File: D, Rank: 5},
				{File: D, Rank: 6},
				{File: D, Rank: 7},
				{File: A, Rank: 3},
				{File: B, Rank: 3},
				{File: C, Rank: 3},
				{File: E, Rank: 3},
				{File: F, Rank: 3},
				{File: G, Rank: 3},
				{File: H, Rank: 3},
			},
		},
	}

	for _, scenario := range tests {
		t.Run(scenario.name, func(t *testing.T) {
			rook := Rook{White}
			got := rook.GetPossibleMoves(scenario.args.position)

			// Check if each element in tt.want is contained in got

			AssertEqualPositions(t, scenario, got)
		})
	}
}
