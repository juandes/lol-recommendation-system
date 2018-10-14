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

// curl -d '{"key1":"value1", "key2":"value2"}' -H "Content-Type: application/json" -X POST http://localhost:8080/recommend
// curl -d '{"champions":["jax"], "intercept": true}' -H "Content-Type: application/json" -X POST http://localhost:8080/recommend

func main() {
	file, err := ioutil.ReadFile("../../data/champions_key.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &champions)

	// read the training set
	train, _, err := data.ReadData("../../data/winning_teams.csv")
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
	}

	log.Info("Starting recommendation service...")

	// create the recommender engine with k = 10
	engine := recommender.NewNeighborhoodBasedRecommender(train, 1)

	mux := http.NewServeMux()

	// endpoints
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/healthz", func(res http.ResponseWriter, _ *http.Request) {
		res.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/recommend", recommendationHandler(engine))

	go log.Fatal(http.ListenAndServe(":8080", mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
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
		}

		for _, val := range pinput.Champions {
			log.Infof("val: %v", val)
			item, ok := champions[val]
			if !ok {
				log.Warningf("Unknown champion: %v", val)
				return
			}
			championIdx, _ := strconv.Atoi(item)
			log.Infof("ChampionIdx: %v", championIdx)
			input[championIdx] = 1
			log.Infof("championIdx: %v", championIdx)
		}

		//log.Infof("Body: %v", string(body))

		recommendation, err := engine.Recommend(input, vectormath.Pearson, pinput.Intercept, pinput.Shuffle, pinput.Serendipity)
		if err != nil {
			log.Errorf("Error predicting recommendation: %v", err)
			return
		}

		for _, recommendation := range recommendation {
			fmt.Println(recommendation.String())
		}
	})
}
