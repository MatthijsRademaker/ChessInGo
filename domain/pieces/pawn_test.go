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
					X: 0,
					Y: 1,
				},
			},
			want: []Position{
				{X: 0, Y: 2},
				{X: 0, Y: 3},
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
				{X: 4, Y: 5},
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
					X: 6,
					Y: 6,
				},
			},
			want: []Position{
				{X: 6, Y: 5},
				{X: 6, Y: 4},
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
				{X: 4, Y: 3},
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
