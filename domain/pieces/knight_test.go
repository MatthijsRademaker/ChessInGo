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
					X: 0,
					Y: 0,
				},
			},
			want: []Position{
				{X: 1, Y: 2},
				{X: 2, Y: 1},
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
				{X: 2, Y: 3},
				{X: 2, Y: 5},
				{X: 3, Y: 2},
				{X: 3, Y: 6},
				{X: 5, Y: 2},
				{X: 5, Y: 6},
				{X: 6, Y: 3},
				{X: 6, Y: 5},
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
