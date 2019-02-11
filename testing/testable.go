package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	params := os.Args[1:]
	sum, err := addNumbers(params...)

	fmt.Println(sum)

	if err != nil {
		os.Exit(1)
	}
}

func addNumbers(nums ...string) (string, error) {
	count := len(nums)

	if count < 2 {
		return "", errors.New("You need to supply at least two numbers")
	}

	var sum float64

	for index, numString := range nums {
		num, err := strconv.ParseFloat(numString, 64)

		if err != nil {
			return "", fmt.Errorf("Number can't be parsed. index: %d; %s", index, numString)
		}

		sum += num
	}

	return strconv.FormatFloat(sum, 'f', -1, 64), nil
}
