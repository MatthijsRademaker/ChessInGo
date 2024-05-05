package pieces

// PieceType represents the type of chess piece
type PieceType int

// Constants representing different types of chess pieces
const (
	PawnType PieceType = iota
	KnightType
	BishopType
	RookType
	QueenType
	KingType
)

func (p PieceType) String() string {
	return [...]string{"Pawn", "Knight", "Bishop", "Rook", "Queen", "King"}[p]
}
