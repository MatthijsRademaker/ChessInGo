package pieces

type Queen struct {
	Color
}

func (p Queen) ToString() string { return "[Q]" }

func (p Queen) GetType() PieceType { return QueenType }

func (p Queen) GetColor() Color { return p.Color }

// GetPossibleMoves calculates possible moves for a Queen
func (p Queen) GetPossibleMoves(position Position) []Position {
	moves := make([]Position, 0)

	// Define all possible diagonal directions (up-left, up-right, down-left, down-right)
	directions := [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}, {-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	// Generate possible moves for each direction
	for _, dir := range directions {
		for i := 1; i < 8; i++ {
			// Max distance queen can move on an 8x8 board
			newX, newY := position.X+i*dir[0], position.Y+i*dir[1]
			// Check if the new position is within the bounds of the chessboard
			if !IsOutOfBounds(newX, newY) {
				// Add the new position to the list of possible moves
				moves = append(moves, Position{newX, newY})
			}
		}
	}

	return moves
}
