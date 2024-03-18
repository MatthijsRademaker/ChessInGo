package pieces

type King struct {
	Color
}

func (p King) ToString() string { return "[G]" }

func (p King) GetType() PieceType { return KingType }

func (p King) GetColor() Color { return p.Color }

// GetPossibleMoves calculates possible moves for a bishop
func (p King) GetPossibleMoves(position Position) []Position {
	moves := make([]Position, 0)

	possibleMoves := [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	// Generate all possible moves from the current position
	for _, move := range possibleMoves {
		newX, newY := move[0], move[1]
		// Check if the new position is within the bounds of the chessboard (assuming an 8x8 board)
		if !IsOutOfBounds(newX, newY) {
			moves = append(moves, Position{newX, newY})
		}
	}

	return moves
}
