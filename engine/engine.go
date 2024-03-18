package main

import (
	"fmt"
	"log"

	"chess/domain"
)

func main() {
	log.SetFlags(0)

	board := domain.InitBoard()

	fmt.Println(board.ToString())

	for !board.IsFinished() {

		// Read a move from the user
		originalPos, destPos, err := ReadMove()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		piece := board.GetPieceAt(originalPos)
		if piece == nil || !board.IsMoveAllowed(piece, originalPos, destPos) {
			fmt.Println("Invalid move try again")
			continue
		}

		board.MovePiece(piece, originalPos, destPos)

		fmt.Print(board.ToString())
	}

	fmt.Println("Congratulations! You won!")
}
