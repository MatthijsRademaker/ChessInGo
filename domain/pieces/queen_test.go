package pieces

import (
	"testing"
)

func TestQueenPossibleMoves(t *testing.T) {
	type args struct {
		position Position
	}
	tests := []struct {
		name string
		args args
		want []Position
	}{
		{
			name: "Test 1",
			args: args{
				position: Position{
					X: 0,
					Y: 0,
				},
			},
			want: []Position{
				{X: 0, Y: 1},
				{X: 0, Y: 2},
				{X: 0, Y: 3},
				{X: 0, Y: 4},
				{X: 0, Y: 5},
				{X: 0, Y: 6},
				{X: 0, Y: 7},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 3, Y: 0},
				{X: 4, Y: 0},
				{X: 5, Y: 0},
				{X: 6, Y: 0},
				{X: 7, Y: 0},
				{X: 1, Y: 1}, // diagonal moves
				{X: 2, Y: 2},
				{X: 3, Y: 3},
				{X: 4, Y: 4},
				{X: 5, Y: 5},
				{X: 6, Y: 6},
				{X: 7, Y: 7},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queen := Queen{White}
			got := queen.GetPossibleMoves(tt.args.position)

			// Check if each element in tt.want is contained in got
			for _, wantMove := range tt.want {
				found := false
				for _, move := range got {
					if move == wantMove {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Queen.GetPossibleMoves() = %v, does not contain %v", got, wantMove)
				}
			}

			// Check if the length of got matches the length of tt.want
			if len(got) != len(tt.want) {
				t.Errorf("Queen.GetPossibleMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
