package quadratic

import (
	"math"
	"testing"
)

func TestSolveQuadratic(t *testing.T) {
	tests := []struct {
		name      string
		a, b, c   float64
		wantX1    float64
		wantX2    float64
		wantError bool
	}{
		{name: "two real roots", a: 1, b: -3, c: 2, wantX1: 2, wantX2: 1},
		{name: "repeated root", a: 1, b: 2, c: 1, wantX1: -1, wantX2: -1},
		{name: "negative discriminant", a: 1, b: 0, c: 1, wantError: true},
		{name: "linear equation", a: 0, b: 2, c: -6, wantX1: 3, wantX2: 3},
		{name: "constant equation", a: 0, b: 0, c: 5, wantError: true},
		{name: "identity", a: 0, b: 0, c: 0, wantError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x1, x2, err := SolveQuadratic(tt.a, tt.b, tt.c)
			if (err != nil) != tt.wantError {
				t.Fatalf("SolveQuadratic(%v, %v, %v) error = %v, wantError %v", tt.a, tt.b, tt.c, err, tt.wantError)
			}
			if err != nil {
				return
			}
			if !closeEnough(x1, tt.wantX1) || !closeEnough(x2, tt.wantX2) {
				t.Errorf("SolveQuadratic(%v, %v, %v) = (%v, %v), want (%v, %v)", tt.a, tt.b, tt.c, x1, x2, tt.wantX1, tt.wantX2)
			}
		})
	}
}

func closeEnough(got, want float64) bool {
	return math.Abs(got-want) <= 1e-9
}
