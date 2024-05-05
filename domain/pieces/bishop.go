package pieces

type Bishop struct {
	Color
}

func (p Bishop) ToString() string { return "[B]" }

func (p Bishop) GetType() PieceType { return BishopType }

func (p Bishop) GetColor() Color { return p.Color }

// GetPossibleMoves calculates possible moves for a bishop
func (p Bishop) GetPossibleMoves(position Position) []Position {
	moves := make([]Position, 0)

	// Define all possible diagonal directions (up-left, up-right, down-left, down-right)
	directions := [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	// Generate possible moves for each direction
	for _, dir := range directions {
		for i := 1; i < 8; i++ { // Max distance bishop can move on an 8x8 board

			x, y := ConvertToGridPosition(position)
			newX, newY := x+i*dir[0], y+i*dir[1]
			// Check if the new position is within the bounds of the chessboard
			if !IsOutOfBounds(newX, newY) {
				// Add the new position to the list of possible moves
				moves = append(moves, ConvertToFileRank(newX, newY))
			}
		}
	}

	return moves
}
