package quadratic

import (
	"math"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
		a, b, c float64
		want []float64
		wantErr bool
	}{
		{name: "two real roots", a: 1, b: -3, c: 2, want: []float64{1, 2}},
		{name: "repeated root", a: 1, b: 2, c: 1, want: []float64{-1}},
		{name: "no real roots", a: 1, b: 0, c: 1, want: []float64{}},
		{name: "linear equation", a: 0, b: 2, c: -6, want: []float64{3}},
		{name: "nonzero constant", a: 0, b: 0, c: 5, want: []float64{}},
		{name: "identity", a: 0, b: 0, c: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Solve(tt.a, tt.b, tt.c)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Solve(%v, %v, %v) error = %v, wantErr %v", tt.a, tt.b, tt.c, err, tt.wantErr)
			}
			if err != nil {
				return
			}
			if !sameRoots(got, tt.want) {
				t.Errorf("Solve(%v, %v, %v) = %v, want %v", tt.a, tt.b, tt.c, got, tt.want)
			}
		})
	}
}

func sameRoots(got, want []float64) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if math.Abs(got[i]-want[i]) > 1e-9 {
			return false
		}
	}
	return true
}
