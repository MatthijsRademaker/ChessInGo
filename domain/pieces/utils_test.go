package pieces

import (
	"testing"
)

func TestIsOutOfBounds(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				x: 0,
				y: 0,
			},
			want: false,
		},
		{
			name: "Test 2",
			args: args{
				x: 7,
				y: 7,
			},
			want: false,
		},
		{
			name: "Test 3",
			args: args{
				x: 8,
				y: 8,
			},
			want: true,
		},
		{
			name: "Test 4",
			args: args{
				x: -1,
				y: -1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOutOfBounds(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("IsOutOfBounds() = %v, want %v", got, tt.want)
			}
		})
	}
}
