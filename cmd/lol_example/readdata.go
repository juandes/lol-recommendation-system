package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readData(path string) ([][]float64, []string, error) {
	var data [][]float64

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// header line
	scanner.Scan()
	header := strings.Split(scanner.Text(), ",")

	for scanner.Scan() {
		var features []float64
		splitFeatures := strings.Split(scanner.Text(), ",")
		for _, val := range splitFeatures {
			feature, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, nil, fmt.Errorf("Error encountered while parsing data file: %v", err)
			}

			features = append(features, feature)
		}
		data = append(data, features)
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data, header, nil
}
