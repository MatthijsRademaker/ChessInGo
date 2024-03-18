package pieces

type Knight struct {
	Color
}

func (p Knight) ToString() string { return "[K]" }

func (p Knight) GetType() PieceType { return KnightType }

func (p Knight) GetColor() Color { return p.Color }

// Todo add boardstate so we can check if the Knight is allowed to move 2 squares or attack diagonally
func (p Knight) GetPossibleMoves(position Position) []Position {
	moves := make([]Position, 0, 8) // Knights have at most 8 possible moves

	// Define all possible knight moves
	possibleMoves := [][]int{
		{-1, -2}, {-2, -1},
		{-2, 1}, {-1, 2},
		{1, 2}, {2, 1},
		{2, -1}, {1, -2},
	}

	// Generate all possible moves from the current position
	for _, move := range possibleMoves {
		newX, newY := position.X+move[0], position.Y+move[1]
		// Check if the new position is within the bounds of the chessboard (assuming an 8x8 board)
		if !IsOutOfBounds(newX, newY) {
			moves = append(moves, Position{newX, newY})
		}
	}

	return moves
}
