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
