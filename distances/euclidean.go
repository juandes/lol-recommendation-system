package distances

import (
	"fmt"
	"math"
)

func EuclideanDistance(x []float64, y []float64) (float64, error) {
	var sum float64
	if len(x) != len(y) {
		return sum, fmt.Errorf("different slice sizes %v, %v", len(x), len(y))
	}

	for index, element := range x {
		sum += math.Pow(element-y[index], 2)
	}

	return math.Sqrt(sum), nil
}
