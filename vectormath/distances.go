package vectormath

import (
	"fmt"
	"github.com/dgryski/go-onlinestats"
	"math"
)

type Distance int

const (
	Euclidean Distance = 0
	Cosine    Distance = 1
	Manhattan Distance = 2
	Pearson   Distance = 3
)

// TODO(Juan): Change the x and y variables to a and b

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

func ManhattanDistance(x []float64, y []float64) (float64, error) {
	var sum float64
	if len(x) != len(y) {
		return sum, fmt.Errorf("different slice sizes len(x): %v, len(y): %v", len(x), len(y))
	}

	for index, element := range x {
		sum += math.Abs(element - y[index])
	}

	return sum, nil
}

// CosineSimilarity calculates the cosine similarity between two vectors.
// cos(d_1, d_2) = (d_1 . d_2) / (||d_1|| * ||d_2||)
func CosineSimilarity(x []float64, y []float64) (float64, error) {
	dot, err := dotProduct(x, y)
	if err != nil {
		return 0, fmt.Errorf("Could not calculate the cosine similarity: %v", err)
	}

	fmt.Printf("cosine: %v\n", float64((dot))/(vectorEuclideanNorm(x)*vectorEuclideanNorm(y)))
	return float64((dot)) / (vectorEuclideanNorm(x) * vectorEuclideanNorm(y)), nil
}

func PearsonCorrelation(x, y []float64) (float64, error) {
	if len(x) != len(y) {
		return 0, fmt.Errorf("different slice sizes len(x): %v, len(y): %v", len(x), len(y))
	}

	// The Pearson correlation coefficient lies between -1 and 1,
	// however I want the score to be from 0 and 1, where
	// 0 represents perfect correlation regardless of whether
	// it is a positive or negative one
	return 1 - math.Abs(onlinestats.Pearson(x, y)), nil
}

// VectorEuclideanNorm calculates the euclidean norm (also known as magnitude or length)
// x = sqrt(x^2_1 + x^2_2 + ... + x^2_n)
func vectorEuclideanNorm(vec []float64) float64 {
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
func dotProduct(x []float64, y []float64) (float64, error) {
	var sum float64
	if len(x) != len(y) {
		return sum, fmt.Errorf("different slice sizes %v, %v", len(x), len(y))
	}

	for i, element := range x {
		sum += element * y[i]
	}

	return sum, nil
}
