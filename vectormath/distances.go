package vectormath

import (
	"fmt"
	"math"
)

type Distance int

const (
	Euclidean Distance = 0
	Cosine    Distance = 1
)

func EuclideanDistance(x []float64, y []float64) (float64, error) {
	var sum float64
	if len(x) != len(y) {
		return sum, fmt.Errorf("different slice sizes len(x): %v, len(y): %v", len(x), len(y))
	}

	for index, element := range x {
		sum += math.Pow(element-y[index], 2)
	}

	return math.Sqrt(sum), nil
}

// CosineSimilarity calculates the cosine similarity between two vectors.
// cos(d_1, d_2) = (d_1 . d_2) / (||d_1|| * ||d_2||)
func CosineSimilarity(x []int64, y []int64) (float64, error) {
	dot, err := dotProduct(x, y)
	if err != nil {
		return 0, fmt.Errorf("Could not calculate the cosine similarity: %v", err)
	}

	return float64((dot)) / (vectorEuclideanNorm(x) * vectorEuclideanNorm(y)), nil
}

// VectorEuclideanNorm calculates the euclidean norm (also known as magnitude or length)
// x = sqrt(x^2_1 + x^2_2 + ... + x^2_n)
func vectorEuclideanNorm(vec []int64) float64 {
	if len(vec) == 0 {
		return 0.0
	}

	var sum float64
	for _, val := range vec {
		sum += math.Pow(float64(val), 2)
	}

	return math.Sqrt(sum)
}

// DotProduct computes the dot product between 2 vectors, d_1 . d_2
func dotProduct(x []int64, y []int64) (int64, error) {
	var sum int64
	if len(x) != len(y) {
		return sum, fmt.Errorf("different slice sizes %v, %v", len(x), len(y))
	}

	for i, element := range x {
		sum += element * y[i]
	}

	return sum, nil
}
