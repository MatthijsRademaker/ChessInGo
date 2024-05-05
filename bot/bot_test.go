package bot

import (
	"chess/domain"
	"chess/domain/pieces"
	"testing"
)

func TestBot_MakeMove(t *testing.T) {
	initBoard := domain.InitBoard()
	generatedBoards := generateExpectedBoards(1, [][]domain.Move{
		{
			{
				From:  pieces.Position{Y: 6, X: 4},
				To:    pieces.Position{Y: 4, X: 4},
				Piece: pieces.Pawn{Color: pieces.White},
			},
		}})

	var expectedBoards []*domain.Board
	for i := range generatedBoards {
		expectedBoards = append(expectedBoards, &generatedBoards[i])
	}

	tests := []TestingScenario{
		{
			name: "Test 1",
			args: &initBoard,
			want: expectedBoards,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot := &Bot{
				Board: &initBoard,
				Color: pieces.White,
			}
			bot.MakeMove()

			AssertEqualBoard(t, tt, bot.Board)
		})
	}
}

func generateExpectedBoards(size int, moves [][]domain.Move) []domain.Board {
	expectedBoards := make([]domain.Board, 0)
	for i := 0; i < size; i++ {
		board := domain.InitBoard()
		for _, move := range moves[i] {
			board.MovePiece(move)
		}
		expectedBoards = append(expectedBoards, board)
	}
	return expectedBoards
}
