package pieces

// ChessPiece represents a chess piece
type ChessPiece interface {
	GetPossibleMoves(Position) []Position
	ToString() string
	GetType() PieceType
	GetColor() Color
}
