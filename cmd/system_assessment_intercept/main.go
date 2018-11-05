package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/juandes/knn-recommender-system/data"
	"github.com/juandes/knn-recommender-system/recommender"
	vm "github.com/juandes/knn-recommender-system/vectormath"
	log "github.com/sirupsen/logrus"
)

func main() {
	var (
		result  []string
		results [][]string
		total   int32
	)

	// read the training set
	train, _, err := data.ReadData("../../data/winning_teams.csv")
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
	}

	// load the testing set; the data that will be used to assess the system
	test, _, err := data.ReadData("../../data/random_teams_large.csv")
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
	}

	for _, i := range []int{1, 2} {
		log.Infof("i: %d", i)
		reco := recommender.NewNeighborhoodBasedRecommender(train, i)

		for _, val := range test {
			recommendations, err := reco.Recommend(val, vm.Pearson, true, true, false)
			if err != nil {
				log.Fatalf("Error while recommending: %v", err)
			}

			result = append(result, strconv.FormatInt(int64(total%4)+1, 10))

			numberItems := 0.0
			for _, val := range recommendations[0].GetRecommendation() {
				numberItems += val
			}

			result = append(result, strconv.FormatInt(int64(numberItems), 10))

			results = append(results, result)
			result = []string{}

			if total%10000 == 0 {
				log.Infof("total: %d", total)
			}
			total++
		}

		file, err := os.Create(fmt.Sprintf("../../data/intercept_test/%v", i))
		if err != nil {
			log.Fatalf("Unable to create file: %v", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		err = writer.WriteAll(results)
		if err != nil {
			log.Fatalf("Unable to create file: %v", err)
		}

		results = [][]string{}
	}

}
