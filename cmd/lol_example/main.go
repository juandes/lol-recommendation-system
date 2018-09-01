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

	data, header, err := readData("winning_teams.csv")
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
	}

	champions := make([]*champion, len(header))

	for i, championName := range header {
		champions[i] = NewChampion(championName)
	}

	v := []float64{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	reco := recommender.NewNeighborhoodBasedRecommender(data, 1)
	recommendations, err := reco.Recommend(v, 3, distances.Euclidean, true)
	if err != nil {
		log.Fatalf("Error while recommending: %v", err)
	}

	fmt.Printf("Recommendation for vector: %v\n", v)
	for _, recommendation := range recommendations {
		fmt.Println(recommendation.String())
	}
}
