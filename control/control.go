package main

import "fmt"

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

	var intarray [20]int

	intarray[0] = 1

	dynamicslice := make([]int, 0, 30)
	second := dynamicslice[15:20]

	//Slice Bestandteile
	// *int startArray
	// int len
	// int cap

	append(second, 5)

	fmt.Println(second)

	return mul, sum
}

func main() {
	mul, sum := whileloop(5)

	fmt.Printf("Our whileloop returns for input 5 mul = %d and sum = %d. \n", mul, sum)
}
