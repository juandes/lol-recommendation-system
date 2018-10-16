package main

import (
	"encoding/json"
	"fmt"
	"github.com/juandes/knn-recommender-system/data"
	"github.com/juandes/knn-recommender-system/recommender"
	"github.com/juandes/knn-recommender-system/vectormath"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
)

var champions map[string]string

// PredictionInput is the object to which the JSON obtain to predict, will be unmarshal to
type PredictionInput struct {
	Champions   []string `json:"champions"`
	Intercept   bool     `json:"intercept"`
	Shuffle     bool     `json:"shuffle"`
	Serendipity bool     `json:"serendipity"`
}

// Output is the return response that holds the recommendations
type Output struct {
	Recommendations []recommender.Recommendation `json:"recommendations"`
}

// TODO (Juan): add an error structure. See this: https://blog.restcase.com/rest-api-error-codes-101/

// curl -d '{"key1":"value1", "key2":"value2"}' -H "Content-Type: application/json" -X POST http://localhost:8080/recommend
// curl -d '{"champions":["jax"], "intercept": true}' -H "Content-Type: application/json" -X POST http://localhost:8080/recommend

func main() {
	// load the file with the mapping of champions name to index position
	file, err := ioutil.ReadFile("../../data/champions_key.json")
	if err != nil {
		log.Fatalf("Error reading champions file: %v", err)
	}
	err = json.Unmarshal(file, &champions)

	// read the training set
	train, _, err := data.ReadData("../../data/winning_teams.csv")
	if err != nil {
		log.Fatalf("Error reading training set: %v", err)
	}

	log.Info("Starting recommendation service...")

	// create the recommender engine with k = 10
	engine := recommender.NewNeighborhoodBasedRecommender(train, 5)

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(res http.ResponseWriter, _ *http.Request) {
		res.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/recommend", recommendationHandler(engine))

	go log.Fatal(http.ListenAndServe(":8080", mux))
}

func recommendationHandler(engine *recommender.NeighborhoodBasedRecommender) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var pinput PredictionInput
		input := make([]float64, len(champions))

		if r.Method != "POST" {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		err = json.Unmarshal(body, &pinput)
		if err != nil {
			log.Errorf("Error unmarshaling: %v", err)
			return
		}

		// create the feature vector by adding 1
		// in the index corresponding to the champion
		for _, val := range pinput.Champions {
			item, ok := champions[val]
			if !ok {
				log.Warningf("Unknown champion: %v", val)
				return
			}
			championIdx, _ := strconv.Atoi(item)
			input[championIdx] = 1
		}

		recommendations, err := engine.Recommend(input, vectormath.Pearson, pinput.Intercept, pinput.Shuffle, pinput.Serendipity)
		if err != nil {
			log.Errorf("Error predicting recommendation: %v", err)
			return
		}

		response, _ := json.Marshal(&Output{
			Recommendations: recommendations,
		})
		fmt.Fprintln(w, string(response))
	})
}
