package main

import "fmt"

// Using the for loop to create a "while" like loop.
func whileloop(runcount int) (mul, sum int) {
	i := 1

	mul = 1
	sum = 0

	// While loops don't exists in go. A for loop without control variable declaration and
	// variable alternation is used. (Left and right part is left out.)
	for true {
		mul *= i
		sum += i

		i++

	}

	return mul, sum
}

// This is an example for the for loop used in it's regular form.
func forloop(runcount int) (mul, sum int) {
	mul = 1
	sum = 0

	// Initializes a control variable i (int); compares with runcount; increments i after one iteration.
	for i := 0; i < runcount; i++ {
		mul *= i
		sum += i
	}

	return mul, sum
}

func main() {
	mul, sum := whileloop(5)

	fmt.Printf("Our whileloop returns for input 5 mul = %d and sum = %d. \n", mul, sum)

	mulfromfor, sumfromfor := forloop(20)

	fmt.Printf("Our for returns for input 20 mul = %d and sum = %d. \n", mulfromfor, sumfromfor)
}
