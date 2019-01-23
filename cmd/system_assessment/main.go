package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"github.com/juandes/lol-recommendation-system/data"
	"github.com/juandes/lol-recommendation-system/recommender"
	vm "github.com/juandes/lol-recommendation-system/vectormath"
	log "github.com/sirupsen/logrus"
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
	train, _, err := data.ReadData("../../static/winning_teams.csv")
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
	}

	// load the testing set; the data that will be used to assess the system
	test, _, err := data.ReadData("../../static/random_teams.csv")
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
	}

	reco := recommender.NewNeighborhoodBasedRecommender(train, 10)

	for _, val := range test {
		total++

		t = time.Now()
		recommendations, err := reco.Recommend(val, vm.Pearson, false, true, false)
		tSince = time.Since(t).Seconds()
		if err != nil {
			log.Fatalf("Error while recommending: %v", err)
		}

		// Append the distance score
		result = append(result, strconv.FormatFloat(recommendations[0].GetDistance(), 'f', 16, 64))
		// Append the time it took to perform the prediction
		result = append(result, strconv.FormatFloat(tSince, 'f', 6, 64))
		results  = append(results, result)
		
		if total%100 == 0 {
			log.Infof("total: %d", total)
		}

		result = []string{}
	}
	writeFile("../../appendix/assessments/distance_time.csv", results)
}


func writeFile(fileName string, data [][]string){
	
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Unable to create file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(data)
	if err != nil {
		log.Fatalf("Unable to create file: %v", err)
	}
}