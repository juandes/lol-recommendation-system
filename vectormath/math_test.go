package vectormath

import (
	"github.com/facebookgo/ensure"
	"testing"
)

func Test_Math(t *testing.T) {
	a := []float64{0, 1, 2, 3, 4}
	b := []float64{5, 6, 7, 8, 9}

	c, err := Intercept(a, b)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, c, []float64{0, 0, 0, 0, 0})

	b = []float64{5, 1, 7, 8, 9}

	c, err = Intercept(a, b)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, c, []float64{0, 1, 0, 0, 0})

	b = a

	c, err = Intercept(a, b)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, c, a)

	b = []float64{}

	_, err = Intercept(a, b)
	ensure.NotNil(t, err)
}
