package domain

type TestingScenario struct {
	name          string
	startingBoard *Board
	move          Move
	want          *Board
}
