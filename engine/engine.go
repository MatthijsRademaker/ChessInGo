package main

import (
	"fmt"
	"log"

	"chess/domain"

	"chess.domain/pieces"
)

func main() {
	log.SetFlags(0)

	board := domain.InitBoard()

	fmt.Println(board.ToString())

	currentPos := pieces.Position{0, 0}

	for !board.IsFinished() {

		piece := board.GetPieceAt(currentPos)
		// Reminder this is bot behavior and should be removed
		moves := board.GetAllowedMoves(piece, currentPos)
		if len(moves) == 0 {
			fmt.Println("No moves left")
			break
		}
		nextPos := moves[0]
		board.MovePiece(piece, currentPos, nextPos)

		fmt.Print(board.ToString())

		currentPos = nextPos
	}

	fmt.Println("Congratulations! You won!")

}
