package main

import (
	"chess/bot"
	"chess/domain"
	"chess/domain/pieces"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(0)

	board := domain.InitBoard()
	bot := bot.Bot{Board: &board, Color: pieces.Black}

	fmt.Println(board.ToString())

	for !board.IsFinished() {

		// Read a move from the user
		originalPos, destPos, err := ReadMove()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		piece := board.GetPieceAt(originalPos)
		move := domain.Move{From: originalPos, To: destPos, Piece: piece}

		if piece == nil || !board.IsMoveAllowed(move) {
			fmt.Println("Invalid move try again")
			continue
		}

		board.MovePiece(move)

		fmt.Print(board.ToString())

		// Make the bot move
		fmt.Println("Bot is thinking...")
		bot.MakeMove()
		fmt.Print(board.ToString())
	}

	fmt.Println("Congratulations! You won!")
}
