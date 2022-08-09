package square

import "math"

type Square struct {
	side float64
}

func NewSquare(side float64) Square {
	if side < 0 {
		panic("side must be non negative")
	}
	return Square{side}
}

func CalculatePerimeter(square Square) float64 {
	if square.side == 0.0 {
		return 0.0
	}
	return square.side + square.side + square.side + square.side
}

func CalculateArea(square Square) float64 {
	if square.side == 0.0 {
		return 0.0
	}
	return math.Round((square.side*square.side)*100000) / 100000
}
