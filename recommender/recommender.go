package recommender

import (
	"fmt"
	"sort"

	log "github.com/Sirupsen/logrus"
	"github.com/juandes/knn-recommender-system/distances"
)

type NeighborhoodBasedRecommender struct {
	data        [][]float64
	neighbors   int
	numberItems int
}

// NewNeighborhoodBasedRecommender creates a new NeighborhoodBasedRecommender object
func NewNeighborhoodBasedRecommender(data [][]float64, neighbors int) *NeighborhoodBasedRecommender {
	if len(data) == 0 {
		log.Fatalf("Dataset is empty")
	}

	return &NeighborhoodBasedRecommender{
		data:        data,
		neighbors:   neighbors,
		numberItems: len(data[0]),
	}
}

// Recommend recommends the n number of items that are closer to a given vector using a given distance measure
func (nbr *NeighborhoodBasedRecommender) Recommend(items []float64, numItemsToRecommend int, distanceMeasure distances.Distance) ([]Recommendation, error) {
	recommendations, err := nbr.findKNearestNeighbors(items, numItemsToRecommend, distanceMeasure)
	if err != nil {
		return nil, fmt.Errorf("Error encountered while finding K nearest neighbors: %v", err)
	}

	return recommendations, nil
}

func (nbr *NeighborhoodBasedRecommender) findKNearestNeighbors(items []float64, n int, distanceMeasure distances.Distance) ([]Recommendation, error) {
	var (
		d                 float64
		err               error
		distancesFromUser []Recommendation
	)

	for i, user := range nbr.data {
		if len(user) != nbr.numberItems {
			return nil, fmt.Errorf("Incorrect number of items in vector")
		}

		switch distanceMeasure {
		case distances.Euclidean:
			d, err = distances.EuclideanDistance(items, user)
		default:
			return nil, fmt.Errorf("Invalid distance measure: %v", distanceMeasure)
		}

		if err != nil {
			return nil, fmt.Errorf("Error calculating distance")
		}

		distancesFromUser = append(distancesFromUser, Recommendation{
			index:    i,
			items:    user,
			d:        d,
			distance: distanceMeasure,
		})
	}

	// sort the recommendations by distance from the given vector
	sort.Slice(distancesFromUser, func(i, j int) bool { return distancesFromUser[i].d < distancesFromUser[j].d })

	return distancesFromUser[:n], nil
}
