package main

import (
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/juandes/knn-recommender-system/champion"
	"github.com/juandes/knn-recommender-system/data"
	"github.com/juandes/knn-recommender-system/recommender"
	vm "github.com/juandes/knn-recommender-system/vectormath"
)

func main() {
	var (
		t      time.Time
		tSince float64
	)
	// one hot encoding of the different items
	// 1 means that the "user" has liked, bought, seen ...
	// the item. 0 means the user has not seen the item.
	data, header, err := data.ReadData("data/winning_teams.csv")
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
	}

	champions := make([]*champion.Champion, len(header))

	for i, championName := range header {
		champions[i] = champion.NewChampion(championName)
	}

	v := []float64{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	t = time.Now()
	reco := recommender.NewNeighborhoodBasedRecommender(data, 1)
	tSince = time.Since(t).Seconds()
	log.Infof("Time (nanoseconds) building model: %f", tSince)

	t = time.Now()
	recommendations, err := reco.Recommend(v, 3, vm.Euclidean, true, true)
	tSince = time.Since(t).Seconds()
	log.Infof("Time (nanoseconds) recommending: %f", tSince)
	if err != nil {
		log.Fatalf("Error while recommending: %v", err)
	}

	fmt.Printf("Recommendation for vector: %v\n", v)
	for _, recommendation := range recommendations {
		fmt.Println(recommendation.String())
	}
}
