package pieces

import (
	"testing"
)

func TestWhitePawnPossibleMoves(t *testing.T) {
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
				{File: A, Rank: 2},
				{File: A, Rank: 3},
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
				{File: D, Rank: 5},
			},
		},
	}

	for _, scenario := range tests {
		t.Run(scenario.name, func(t *testing.T) {
			pawn := Pawn{White}
			got := pawn.GetPossibleMoves(scenario.args.position)

			// Check if each element in tt.want is contained in got
			AssertEqualPositions(t, scenario, got)
		})
	}
}

func TestBlackPawnPossibleMoves(t *testing.T) {
	tests := [2]TestingScenario{
		{
			name: "Test 1",
			args: TestingArgs{
				position: Position{
					File: F,
					Rank: 6,
				},
			},
			want: []Position{
				{File: F, Rank: 5},
				{File: F, Rank: 4},
			},
		},
		{
			name: "Test 2",
			args: TestingArgs{
				position: Position{
					File: E,
					Rank: 4,
				},
			},
			want: []Position{
				{File: E, Rank: 3},
			},
		},
	}

	for _, scenario := range tests {
		t.Run(scenario.name, func(t *testing.T) {
			pawn := Pawn{Black}
			got := pawn.GetPossibleMoves(scenario.args.position)

			// Check if each element in tt.want is contained in got
			AssertEqualPositions(t, scenario, got)
		})
	}
}
