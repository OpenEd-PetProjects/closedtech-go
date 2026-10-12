package quadratic

import (
	"errors"
	"math"
)

// Solve returns the real roots of ax² + bx + c = 0 in ascending order.
// A constant nonzero equation has no roots. The equation 0 = 0 has
// infinitely many roots and is reported as an error.
func Solve(a, b, c float64) ([]float64, error) {
	if a == 0 {
		if b == 0 {
			if c == 0 {
				return nil, errors.New("equation has infinitely many solutions")
			}
			return []float64{}, nil
		}
		return []float64{-c / b}, nil
	}

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return []float64{}, nil
	}
	if discriminant == 0 {
		return []float64{-b / (2 * a)}, nil
	}

	root := math.Sqrt(discriminant)
	first := (-b - root) / (2 * a)
	second := (-b + root) / (2 * a)
	if first > second {
		first, second = second, first
	}
	return []float64{first, second}, nil
}
