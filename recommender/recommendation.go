package recommender

import (
	"fmt"

	vm "github.com/juandes/knn-recommender-system/vectormath"
)

type Recommendation interface {
	String() string
	Distance() float64
	Items() []float64
}

type MultipleRecommendation struct {
	index    int
	items    []float64
	d        float64
	distance vm.Distance
}

type SimpleRecommendation struct {
	item     []float64
	distance vm.Distance
}

func (r MultipleRecommendation) String() string {
	return fmt.Sprintf("Items: %v\nIndex: %d\nDistance (%v): %f\n", r.items, r.index, r.distance, r.d)
}

func (r MultipleRecommendation) Distance() float64 {
	return r.d
}

func (r MultipleRecommendation) Items() []float64 {
	return r.items
}

func (r SimpleRecommendation) String() string {
	return fmt.Sprintf("Items: %v\nDistance used: %v\n", r.item, r.distance)
}

func (r SimpleRecommendation) Distance() float64 {
	return 0.0
}

func (r SimpleRecommendation) Items() []float64 {
	return r.item
}
