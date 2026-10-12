package quadratic

import (
	"errors"
	"math"
)

// SolveQuadratic returns the roots of ax² + bx + c = 0.
// For a linear equation, the single root is returned in both root positions.
func SolveQuadratic(a, b, c float64) (float64, float64, error) {
	if a == 0 {
		if b == 0 {
			return 0, 0, errors.New("equation has no unique finite set of roots")
		}
		root := -c / b
		return root, root, nil
	}

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return 0, 0, errors.New("equation has no real roots")
	}
	root := math.Sqrt(discriminant)
	x1 := (-b + root) / (2 * a)
	x2 := (-b - root) / (2 * a)
	return x1, x2, nil
}
