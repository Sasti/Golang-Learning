package main

import (
	"fmt"
	"math"
)

func main() {

	var number = 42

	var number32 int32 = 32
	fmt.Printf("Die Antwort den Sinn des Lebens %d !!! \n", number)

	var number32Pointer *int32

	number32Pointer = &number32

	number32pointer := &number32

	*number32Pointer = 42

	fmt.Printf("Die Antwort den Sinn des Lebens %d !!! \n", *number32pointer)

	var pi32 float32 = math.Pi
	var etwaPi = int(pi32)

	var unsignedzahl uint32 = 0x00000101
	var unsignedzahl2 uint32 = 0x10000110

	and := unsignedzahl ^ unsignedzahl2

	fmt.Printf("Beide Variablen mit UND verknüft %X \n", and)

	fmt.Printf("Die Konstante Pi %f und als float32 %f \n", math.Pi, pi32)
	fmt.Printf("Aus der Schule wissen wir Pi is etwa %d. \n", etwaPi)
	fmt.Println("Woher kommt der Unterschied???")

	*number32Pointer = 3

	exampleString := "Babymarkt"

	exampleStringPointer := &exampleString

	*exampleStringPointer += " lernt golang."

	fmt.Println(*exampleStringPointer)
}
