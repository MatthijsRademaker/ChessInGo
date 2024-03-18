package pieces

type Pawn struct {
	Color
}

func (p Pawn) ToString() string { return "[P]" }

func (p Pawn) GetType() PieceType { return PawnType }

func (p Pawn) GetColor() Color { return p.Color }

// GetPossibleMoves calculates possible moves for a pawn
func (p Pawn) GetPossibleMoves(position Position) []Position {
	moves := make([]Position, 0)

	// Define direction based on pawn's color (up for white, down for black)
	direction := 1 // Assuming white pawns start at the bottom of the board

	// Adjust direction for black pawns
	if p.Color == Black {
		direction = -1
	}

	// Generate possible moves for a single step forward
	moves = append(moves, Position{position.X, position.Y + direction})

	// If the pawn is at its starting position, it can move two squares forward
	if (position.Y == 1 && p.Color == White) || (position.Y == 6 && p.Color == Black) {
		moves = append(moves, Position{position.X, position.Y + 2*direction})
	}

	// TODO: implement logic for en passant and Diagonal attack moves, should be in the boardstate? Or separate function?
	diagonalMoves := []Position{{position.X - 1, position.Y + direction}, {position.X + 1, position.Y + direction}}
	for _, diagonalMove := range diagonalMoves {
		if !IsOutOfBounds(diagonalMove.X, diagonalMove.Y) {
			moves = append(moves, diagonalMove)
		}
	}

	return moves
}
