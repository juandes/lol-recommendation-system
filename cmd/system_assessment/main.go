package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/juandes/knn-recommender-system/data"
	"github.com/juandes/knn-recommender-system/recommender"
	vm "github.com/juandes/knn-recommender-system/vectormath"
)

func main() {
	var (
		t       time.Time
		tSince  float64
		result  []string
		results [][]string
		total   int32
	)

	// read the training set
	train, _, err := data.ReadData("data/winning_teams.csv")
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
	}

	// load the testing set; the data that will be used to assess the system
	test, _, err := data.ReadData("data/random_teams.csv")
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
	}

	// create the recommender engine with k = 10
	reco := recommender.NewNeighborhoodBasedRecommender(train, 10)

	for _, val := range test {
		total++

		t = time.Now()
		recommendations, err := reco.Recommend(val, 1, vm.Pearson, false, true, false)
		tSince = time.Since(t).Seconds()
		if err != nil {
			log.Fatalf("Error while recommending: %v", err)
		}

		// Append the id of the prediction
		// These values had to be changed to string, because Golang's csv library deals only with string
		//result = append(result, string(total))
		// Append the distance score
		result = append(result, strconv.FormatFloat(recommendations[0].Distance(), 'f', 16, 64))
		// Append the time it took to perform the prediction
		result = append(result, strconv.FormatFloat(tSince, 'f', 6, 64))
		/* In the testing set, each 4 rows represent one team, where the first row is the team with
		only one champion, the second row is a team composition of two champions and so on.
		Since in this example, we are analyzing the decrease of distance as more heroes are being
		added to the team, I want to keep the distances related to one team in the same row,
		that's why I am appending it to the results dataset every 4 iterations*/
		if total%4 == 0 {
			results = append(results, result)
			fmt.Printf("result: %v\n", result)
			result = []string{}

		}

		if total%100 == 0 {
			log.Infof("total: %d", total)
		}
	}

	file, err := os.Create("data/assessment_02Pearson_Shuffle.csv")
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

}
