package pieces

func IsOutOfBounds(x, y int) bool {
	return x < 0 || x > 7 || y < 0 || y > 7
}

func (p *Position) IsOutOfBounds() bool {
	x, y := ConvertToGridPosition(*p)
	return IsOutOfBounds(x, y)
}
