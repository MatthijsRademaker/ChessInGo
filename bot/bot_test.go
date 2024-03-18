package bot

import (
	"chess/domain"
	"chess/domain/pieces"
	"testing"
)

func TestBot_MakeMove(t *testing.T) {
	initBoard := domain.InitBoard()
	tests := []TestingScenario{
		{
			name: "Test 1",
			args: &initBoard,
			want: &expectedBoard,
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

func generateExpectedBoards(size int) []domain.Board {
	expectedBoards := make([]domain.Board, 0)
	for i := 0; i < size; i++ {
		expectedBoard := domain.InitBoard()
		piece := expectedBoard.GetPieceAt(pieces.Position{X: 0, Y: 1})
		ep.Move(pieces.Position{X: 0, Y: 3})
		expectedBoards = append(expectedBoards, domain.InitBoard())
	}

}
