package domain

import (
	"chess/domain/pieces"
)

type Move struct {
	From  pieces.Position
	To    pieces.Position
	Piece pieces.ChessPiece
}

// arrays are indexed [row][column] therefore the first index is the Y coordinate and the second index is the X coordinate
func (b *Board) MovePiece(move Move) {
	b.State[move.To.Y][move.To.X] = move.Piece
	b.State[move.From.Y][move.From.X] = nil
}

func (b *Board) GetPieceAt(position pieces.Position) pieces.ChessPiece {
	return b.State[position.Y][position.X]
}

func (b *Board) IsFinished() bool {
	finalPlace := b.State[7][0]
	return finalPlace != nil && finalPlace.GetType() == pieces.RookType && finalPlace.GetColor() == pieces.White
}

func (b *Board) GetState() [8][8]pieces.ChessPiece {
	return b.State
}

func (b *Board) IsMoveAllowed(moveToMake Move) bool {
	allowedMoves := b.GetAllowedMoves(moveToMake.Piece, moveToMake.From)
	for _, move := range allowedMoves {
		if move == moveToMake.To {
			return true
		}
	}
	return false
}

func (b *Board) GetAllowedMoves(piece pieces.ChessPiece, from pieces.Position) []pieces.Position {
	allowedMoves := make([]pieces.Position, 0)
	possibleMoves := piece.GetPossibleMoves(from)
	for _, move := range possibleMoves {
		pieceAtMove := b.GetPieceAt(move)
		if pieceAtMove == nil || pieceAtMove.GetColor() != piece.GetColor() {
			allowedMoves = append(allowedMoves, move)
		}
	}
	return allowedMoves
}
