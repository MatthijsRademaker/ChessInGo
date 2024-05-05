package domain

import (
	"chess/domain/pieces"
	"reflect"
	"testing"
)

func Test_MakeMove(t *testing.T) {
	initBoard := InitBoard()
	expectedBoard := Board{
		State: [8][8]pieces.ChessPiece{
			{&pieces.Rook{pieces.White}, &pieces.Knight{pieces.White}, &pieces.Bishop{pieces.White}, &pieces.Queen{pieces.White}, &pieces.King{pieces.White}, &pieces.Bishop{pieces.White}, &pieces.Knight{pieces.White}, &pieces.Rook{pieces.White}},
			{&pieces.Pawn{pieces.White}, &pieces.Pawn{pieces.White}, &pieces.Pawn{pieces.White}, &pieces.Pawn{pieces.White}, &pieces.Pawn{pieces.White}, &pieces.Pawn{pieces.White}, &pieces.Pawn{pieces.White}, &pieces.Pawn{pieces.White}},
			{nil, nil, nil, nil, nil, nil, nil, nil},
			{nil, nil, nil, nil, nil, nil, nil, nil},
			{nil, nil, nil, &pieces.Pawn{pieces.Black}, nil, nil, nil, nil},
			{nil, nil, nil, nil, nil, nil, nil, nil},
			{&pieces.Pawn{pieces.Black}, &pieces.Pawn{pieces.Black}, &pieces.Pawn{pieces.Black}, nil, &pieces.Pawn{pieces.Black}, &pieces.Pawn{pieces.Black}, &pieces.Pawn{pieces.Black}, &pieces.Pawn{pieces.Black}},
			{&pieces.Rook{pieces.Black}, &pieces.Knight{pieces.Black}, &pieces.Bishop{pieces.Black}, &pieces.Queen{pieces.Black}, &pieces.King{pieces.Black}, &pieces.Bishop{pieces.Black}, &pieces.Knight{pieces.Black}, &pieces.Rook{pieces.Black}},
		},
	}

	tests := []TestingScenario{
		{
			name:          "Test 1",
			startingBoard: &initBoard,
			move: Move{
				From:  pieces.Position{Y: 6, X: 4},
				To:    pieces.Position{Y: 4, X: 4},
				Piece: pieces.Pawn{Color: pieces.Black},
			},
			want: &expectedBoard,
		},
	}
	for _, testScenario := range tests {
		t.Run(testScenario.name, func(t *testing.T) {
			wantedBoard := testScenario.want.GetState()
			initBoard.MovePiece(testScenario.move)

			gotBoard := initBoard.GetState()

			t.Log(initBoard.ToString())
			t.Log(testScenario.want.ToString())
			if !reflect.DeepEqual(wantedBoard, gotBoard) {
				t.Errorf("Test %s failed.", testScenario.name)
			}
		})
	}
}
