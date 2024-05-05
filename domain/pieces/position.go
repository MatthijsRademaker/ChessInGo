package pieces

import "fmt"

type Position struct {
	File File
	Rank int
}

type File int

const (
	A File = iota
	B
	C
	D
	E
	F
	G
	H
)

// NewPosition creates a new position
// file should be a lowercase letter from 'a' to 'h'
func NewPosition(file File, rank int) (*Position, error) {
	if file < 'a' || file > 'h' {
		return nil, fmt.Errorf("invalid file: %c, file should be in the range 'a' to 'h'", file)
	}
	if rank < 1 || rank > 8 {
		return nil, fmt.Errorf("invalid rank: %d, rank should be in the range 1 to 8", rank)
	}
	return &Position{File: file, Rank: rank}, nil
}

// ConvertToGridPosition converts a position to grid coordinates
// (File, Rank) -> (x, y)
func ConvertToGridPosition(position Position) (int, int) {
	return int(position.File), position.Rank - 1
}

func GetFileFromGridPosition(x int) File {
	// TODO does this work?
	return File(x)
}

func GetRankFromGridPosition(y int) int {
	return y
}

func ConvertToFileRank(x, y int) Position {
	return Position{File: GetFileFromGridPosition(x), Rank: GetRankFromGridPosition(y)}
}
