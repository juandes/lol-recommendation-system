package recommender

import (
	"fmt"

	vm "github.com/juandes/knn-recommender-system/vectormath"
)

type Recommendation interface {
	String() string
	GetDistance() float64
	GetRecommendation() []float64
}

type MultipleRecommendation struct {
	Index           int         `json:"-"`
	Recommendation  []float64   `json:"recommendation"`
	Distance        float64     `json:"distance"`
	DistanceMeasure vm.Distance `json:"-"`
}

type SingleRecommendation struct {
	Recommendation  []float64   `json:"recommendation"`
	DistanceMeasure vm.Distance `json:"-"`
}

type SerendipitousRecommendation struct {
	Recommendation  []float64   `json:"recommendation"`
	DistanceMeasure vm.Distance `json:"-"`
}

func (r MultipleRecommendation) String() string {
	return fmt.Sprintf("Items: %v\nIndex: %d\nDistance (%v): %f\n", r.Recommendation, r.Index, r.DistanceMeasure, r.Distance)
}

func (r MultipleRecommendation) GetDistance() float64 {
	return r.Distance
}

func (r MultipleRecommendation) GetRecommendation() []float64 {
	return r.Recommendation
}

func (r SingleRecommendation) String() string {
	return fmt.Sprintf("Items: %v\nDistance used: %v\n", r.Recommendation, r.DistanceMeasure)
}

func (r SingleRecommendation) GetDistance() float64 {
	return 0.0
}

func (r SingleRecommendation) GetRecommendation() []float64 {
	return r.Recommendation
}

func (r SerendipitousRecommendation) String() string {
	return fmt.Sprintf("Serendipitous recommendation items: %v\nDistance used: %v\n", r.Recommendation, r.DistanceMeasure)
}

func (r SerendipitousRecommendation) GetDistance() float64 {
	return 0.0
}

func (r SerendipitousRecommendation) GetRecommendation() []float64 {
	return r.Recommendation
}
