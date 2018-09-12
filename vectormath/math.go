package vectormath

import (
	"fmt"
)

func sum(a []float64, b []float64) ([]float64, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("Length mismatched; len(a)=%d, len(b)=%d", len(a), len(b))
	}
	x := make([]float64, len(a))

	for idx, val := range a {
		x[idx] = val + b[idx]
	}

	return x, nil
}
