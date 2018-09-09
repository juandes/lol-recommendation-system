package recommender

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	log "github.com/Sirupsen/logrus"
	vm "github.com/juandes/knn-recommender-system/vectormath"
)

type NeighborhoodBasedRecommender struct {
	data        [][]float64
	neighbors   int
	numberItems int
}

type Slice struct {
	sort.Interface
	idx []int
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
func (nbr *NeighborhoodBasedRecommender) Recommend(items []float64, numItemsToRecommend int, distanceMeasure vm.Distance, shuffle bool, serendipitous bool) ([]Recommendation, error) {
	recommendations, err := nbr.findKNearestNeighbors(items, numItemsToRecommend, distanceMeasure, shuffle, serendipitous)
	if err != nil {
		return nil, fmt.Errorf("Error encountered while finding K nearest neighbors: %v", err)
	}

	return recommendations, nil
}

func (nbr *NeighborhoodBasedRecommender) findKNearestNeighbors(items []float64, n int, distanceMeasure vm.Distance, shuffle bool, serendipitous bool) ([]Recommendation, error) {
	var (
		d                 float64
		err               error
		distancesFromUser []Recommendation
		recommendations   []Recommendation
		order             []int
	)

	order = make([]int, len(nbr.data))
	for i := range order {
		order[i] = i
	}

	// the point of shuffling the order in which
	// the distances will be calculated
	// is to avoid having always the same
	// predictions in case all the n results
	// to return have the same distance
	if shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(order), func(i, j int) { order[i], order[j] = order[j], order[i] })
	}

	for i, val := range order {
		user := nbr.data[val]
		if len(user) != nbr.numberItems {
			return nil, fmt.Errorf("Incorrect number of items in vector")
		}

		switch distanceMeasure {
		case vm.Euclidean:
			d, err = vm.EuclideanDistance(items, user)
		default:
			return nil, fmt.Errorf("Invalid distance measure: %v", distanceMeasure)
		}

		if err != nil {
			return nil, fmt.Errorf("Error calculating distance: %v", err)
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
	recommendations = distancesFromUser[:n]

	// The idea here is the following:
	// 1. Get the n:n*2 neighbors
	// 2. Build a map where the keys are the champions found on those neighbors
	//    and value is the count of them.
	// 3. Sort the map by its value
	// 4. Use the 5 champions with the highest count as a recommendation
	if serendipitous {
		sereOptions := make([]int, len(nbr.data))
		for _, dist := range distancesFromUser[n : n*2] {
			for j := range dist.items {
				sereOptions[j]++
			}
		}
		//log.Printf("sereOptions: %v", len(sereOptions))
		s := NewIntSlice(sereOptions)
		sort.Sort(s)
		//log.Printf("sere: %v", s.idx[1:5])
	}

	return recommendations, nil
}

func NewSlice(n sort.Interface) *Slice {
	s := &Slice{Interface: n, idx: make([]int, n.Len())}
	for i := range s.idx {
		s.idx[i] = i
	}
	return s
}

func NewIntSlice(n []int) *Slice { return NewSlice(sort.IntSlice(n)) }
