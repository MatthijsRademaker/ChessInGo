package bot

import (
	"chess/domain"
	"chess/domain/pieces"
	"cmp"
	"slices"
)

type Bot struct {
	Board *domain.Board
	Color pieces.Color
}

func (bot *Bot) SimulateMove(move domain.Move) {

	//TODO
}

func (bot *Bot) MakeMove() {
	bestMove := bot.getBestMove()
	bot.Board.MovePiece(domain.Move{Piece: bestMove.Piece, From: bestMove.From, To: bestMove.To})
}

func (bot *Bot) getBestMove() MoveWithPoints {
	currentState := bot.Board.GetState()
	pointsMap := generatePointsMap()

	movesWithPointsSlice := make([]MoveWithPoints, 0)
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			piece := currentState[y][x]
			if piece != nil && piece.GetColor() == bot.Color {
				allowedMoves := bot.Board.GetAllowedMoves(piece, pieces.Position{X: x, Y: y})
				for _, move := range allowedMoves {
					movesWithPointsSlice = append(movesWithPointsSlice, MoveWithPoints{Piece: piece, From: pieces.Position{X: x, Y: y}, To: move, Score: pointsMap[move]})
				}
			}
		}
	}

	bestMoveWithPoint := slices.MaxFunc(movesWithPointsSlice, func(a, b MoveWithPoints) int {
		return cmp.Compare(a.Score, b.Score)
	})

	return bestMoveWithPoint
}
