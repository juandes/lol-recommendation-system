package main

import (
	"fmt"
	"time"

	"github.com/juandes/knn-recommender-system/data"
	"github.com/juandes/knn-recommender-system/recommender"
	vm "github.com/juandes/knn-recommender-system/vectormath"
	log "github.com/sirupsen/logrus"
)

func main() {
	var (
		t      time.Time
		tSince float64
	)
	// one hot encoding of the different items
	// 1 means that the "user" has liked, bought, seen ...
	// the item. 0 means the user has not seen the item.
	data, _, err := data.ReadData("../../data/winning_teams.csv")
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
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
	recommendations, err := reco.Recommend(v, vm.Pearson, true, true, false)
	if err != nil {
		log.Fatalf("Error while recommending: %v", err)
	}

	tSince = time.Since(t).Seconds()
	log.Infof("Time (nanoseconds) recommending: %f", tSince)

	for _, recommendation := range recommendations {
		fmt.Println(recommendation.String())
	}
}
