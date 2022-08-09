package square

type square struct {
	side float64
}

func NewSquare(side float64) square {
	if side < 0 {
		panic("side must be non negative")
	}
	return square{side}
}

func CalculatePerimeter(Square square) float64 {
	if Square.side == 0.0 {
		return 0.0
	}
	return Square.side + Square.side + Square.side + Square.side
}
