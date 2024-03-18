package domain

import (
	"chess.domain/pieces"
)

// arrays are indexed [row][column] therefore the first index is the Y coordinate and the second index is the X coordinate
func (b *Board) MovePiece(piece pieces.ChessPiece, from, to pieces.Position) {
	b.State[to.Y][to.X] = piece
	b.State[from.Y][from.X] = nil
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
