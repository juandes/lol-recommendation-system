package recommender

import (
	"fmt"

	"github.com/juandes/knn-recommender-system/distances"
)

type Recommendation struct {
	index    int
	items    []float64
	d        float64
	distance distances.Distance
}

func (r *Recommendation) String() string {
	return fmt.Sprintf("Items: %v\nIndex: %d\nDistance (%v): %f", r.items, r.index, r.distance, r.d)
}
