package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
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

	var (
		lineNum  int
		valueSum float64
	)

	{ // Print all data and scavange necessary data for further processing
		for {
			data, err := csvReader.Read()
			if err == io.EOF {
				// No more data to read
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			// Print current line
			fmt.Println(data)

			{ // Scavange data
				if lineNum >= 1 { // The first row contains no data, it contains the header so we wont evaluate it
					value, _ := strconv.ParseFloat(data[8], 32)
					valueSum += value
				}
			}

			// Increment line counter
			lineNum++
		}
	}

	// Calculate resutls
	avarageTemp := valueSum / float64(lineNum)

	{ // Export calculated results as csv file
		if resultFile, err := os.Create("result.csv"); err == nil {
			resultWriter := csv.NewWriter(resultFile)
			defer resultFile.Close()
			defer resultWriter.Flush()

			header := []string{
				"element count", "temp sum", "temp average",
			}
			resultWriter.Write(header)

			resultline := []string{
				strconv.FormatInt(int64(lineNum), 10),
				strconv.FormatFloat(valueSum, 'f', -1, 32),
				strconv.FormatFloat(avarageTemp, 'f', -1, 32),
			}
			resultWriter.Write(resultline)
		} else {
			log.Fatalln(err)
		}
	}
}
