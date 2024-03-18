package domain

import (
	"bytes"
	"fmt"

	"chess/domain/pieces"
)

type Board struct {
	State [8][8]pieces.ChessPiece
}

func InitBoard() Board {
	board := Board{
		State: [8][8]pieces.ChessPiece{
			{&pieces.Rook{pieces.White}, &pieces.Knight{pieces.White}, &pieces.Bishop{pieces.White}, &pieces.Queen{pieces.White}, &pieces.King{pieces.White}, &pieces.Bishop{pieces.White}, &pieces.Knight{pieces.White}, &pieces.Rook{pieces.White}},
			{&pieces.Pawn{pieces.White}, &pieces.Pawn{pieces.White}, &pieces.Pawn{pieces.White}, &pieces.Pawn{pieces.White}, &pieces.Pawn{pieces.White}, &pieces.Pawn{pieces.White}, &pieces.Pawn{pieces.White}, &pieces.Pawn{pieces.White}},
			{nil, nil, nil, nil, nil, nil, nil, nil},
			{nil, nil, nil, nil, nil, nil, nil, nil},
			{nil, nil, nil, nil, nil, nil, nil, nil},
			{nil, nil, nil, nil, nil, nil, nil, nil},
			{&pieces.Pawn{pieces.Black}, &pieces.Pawn{pieces.Black}, &pieces.Pawn{pieces.Black}, &pieces.Pawn{pieces.Black}, &pieces.Pawn{pieces.Black}, &pieces.Pawn{pieces.Black}, &pieces.Pawn{pieces.Black}, &pieces.Pawn{pieces.Black}},
			{&pieces.Rook{pieces.Black}, &pieces.Knight{pieces.Black}, &pieces.Bishop{pieces.Black}, &pieces.Queen{pieces.Black}, &pieces.King{pieces.Black}, &pieces.Bishop{pieces.Black}, &pieces.Knight{pieces.Black}, &pieces.Rook{pieces.Black}},
		},
	}

	return board
}
func (b *Board) ToString() string {
	var buffer bytes.Buffer
	buffer.WriteString("+-----------------------------+\n")
	buffer.WriteString("     A  B  C  D  E  F  G  H\n")
	for x := 0; x < 8; x++ {
		buffer.WriteString("  ")

		buffer.WriteString(fmt.Sprintf("%d ", x+1))

		for y := 0; y < 8; y++ {
			piece := b.State[x][y]
			if piece == nil {
				buffer.WriteString("[ ]")
			} else {
				buffer.WriteString(piece.ToString())
			}
		}
		buffer.WriteString("\n")
	}
	buffer.WriteString("     A  B  C  D  E  F  G  H\n")
	buffer.WriteString("+-----------------------------+\n")
	return buffer.String()
}
