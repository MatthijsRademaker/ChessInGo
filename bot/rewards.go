package bot

import (
	"chess/domain/pieces"
	"math"
)

type MoveWithPoints struct {
	Piece pieces.ChessPiece
	From  pieces.Position
	To    pieces.Position
	Score int
}

const (
	BoardSize    = 8
	CenterWeight = 20
	EdgeWeight   = 10
	CornerWeight = 5
)

func generatePointsMap() map[pieces.Position]int {
	pointsMap := make(map[pieces.Position]int)

	// Calculate the center of the board
	centerX := (BoardSize - 1) / 2
	centerY := (BoardSize - 1) / 2

	// Assign points to each position based on their distance from the center
	for x := 0; x < BoardSize; x++ {
		for y := 0; y < BoardSize; y++ {
			distance := math.Sqrt(float64((x-centerX)*(x-centerX) + (y-centerY)*(y-centerY)))
			move := pieces.ConvertToFileRank(x, y)
			switch {
			case x == 0 || x == BoardSize-1 || y == 0 || y == BoardSize-1:
				pointsMap[move] = CornerWeight
			case x == 1 || x == BoardSize-2 || y == 1 || y == BoardSize-2:
				pointsMap[move] = EdgeWeight
			default:
				pointsMap[move] = CenterWeight - int(distance)
			}
		}
	}

	return pointsMap
}
