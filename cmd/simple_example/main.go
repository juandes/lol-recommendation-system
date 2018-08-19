package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/juandes/knn-recommender-system/distances"
	"github.com/juandes/knn-recommender-system/recommender"
)

func main() {
	// one hot encoding of the different items
	// 1 means that the "user" has liked, bought, seen ...
	// the item. 0 means the user has not seen the item.
	data := [][]float64{
		[]float64{1.0, 1.0, 1.0, 0.0, 1.0, 0.0},
		[]float64{1.0, 1.0, 0.0, 0.0, 1.0, 0.0},
		[]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
		[]float64{0.0, 0.0, 0.0, 0.0, 1.0, 0.0},
	}

	reco := recommender.NewNeighborhoodBasedRecommender(data, 1)
	recommendations, err := reco.Recommend([]float64{0.0, 0.0, 0.0, 0.0, 1.0, 0.0}, 1, distances.Euclidean)
	if err != nil {
		log.Fatalf("Error while recommending: %v", err)
		return
	}

	fmt.Printf("Recommended items: %v", recommendations)
	fmt.Printf("Recommended items\n")
	for _, recommendation := range recommendations {
		fmt.Printf(recommendation.String())
	}
}
