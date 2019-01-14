package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Explains the data that is going to be processed
	description, err := http.Get("https://archive.ics.uci.edu/ml/machine-learning-databases/forest-fires/forestfires.names")

	if err != nil {
		log.Fatalf("Could not load description %s", err)
	}

	readme, _ := ioutil.ReadAll(description.Body)
	fmt.Println(string(readme))

	// https://archive.ics.uci.edu/ml/datasets/Forest+Fires
	response, err := http.Get("https://archive.ics.uci.edu/ml/machine-learning-databases/forest-fires/forestfires.csv")

	if err != nil {
		log.Fatalf("Could not load data %s", err)
	}

	// Create a csv Reader to process the data
	csvReader := csv.NewReader(response.Body)

	// Print all data
	for {
		data, err := csvReader.Read()
		if err == io.EOF {
			// No more data to read
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(data)
	}
}
