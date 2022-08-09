package square

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSquare(t *testing.T) {
	t.Run("should create square object", func(t *testing.T) {
		assert.IsType(t, Square{}, NewSquare(5))
	})

	t.Run("should raise exception when side is negative", func(t *testing.T) {
		assert.Panics(t, func() {
			NewSquare(-2.5)
		})
	})

	t.Run("should not raise exception when side is non negative", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewSquare(5.76)
		})
	})
}

func TestCalculatePerimeter(t *testing.T) {
	t.Run("should return Perimeter as zero when side is zero", func(t *testing.T) {
		assert.Equal(t, 0.0, CalculatePerimeter(NewSquare(0.0)))
	})

	t.Run("should return Perimeter as four times of side when side is of length greater than zero", func(t *testing.T) {
		assert.Equal(t, 4.0, CalculatePerimeter(NewSquare(1.0)))
	})

	t.Run("should calculate Perimeter correctly when integers are passed", func(t *testing.T) {
		assert.Equal(t, 8.0, CalculatePerimeter(NewSquare(2)))
	})

	t.Run("should calculate Perimeter correctly when side length is of long float numbers", func(t *testing.T) {
		assert.Equal(t, 2617.3027156, CalculatePerimeter(NewSquare(654.3256789)))
	})
}

func TestCalculateArea(t *testing.T) {
	t.Run("should return Area as zero when side is zero", func(t *testing.T) {
		assert.Equal(t, 0.0, CalculateArea(NewSquare(0)))
	})

	t.Run("should return Area as square of side when side length is greater than zero", func(t *testing.T) {
		assert.Equal(t, 2.56, CalculateArea(NewSquare(1.60)))
	})

	t.Run("should return Area rounded of to five decimal places when side length is of a fraction", func(t *testing.T) {
		assert.Equal(t, 572497.87004, CalculateArea(NewSquare(756.63589)))
	})

	t.Run("should calculate Area correctly when side is of an integer type", func(t *testing.T) {
		assert.Equal(t, 25.0, CalculateArea(NewSquare(5)))
	})
}
