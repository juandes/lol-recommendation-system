package vectormath

import (
	"fmt"
)

func sum(a []float64, b []float64) ([]float64, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("Length mismatched; len(a)=%d, len(b)=%d", len(a), len(b))
	}
	c := make([]float64, len(a))

	for idx, val := range a {
		c[idx] = val + b[idx]
	}

	return c, nil
}

func Intercept(a []float64, b []float64) ([]float64, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("Length mismatched; len(a)=%d, len(b)=%d", len(a), len(b))
	}
	c := make([]float64, len(a))

	for idx, val := range a {
		if val == b[idx] {
			c[idx] = val
		}
	}

	return c, nil
}
