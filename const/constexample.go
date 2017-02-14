package main

import (
	"fmt"
	"net/http"
)

	const (
		Montag Wochentag = 1 << iota
		Dienstag
		Mittwoch
		Donnerstag
		Freitag
		Samstag
		Sonntag
	)

func main() {
	const pi = 3.14

	const x = 2890284901489018348231948902183904892134
	const y = 3

	var laufzeit = float64(x + y)

	fmt.Println(laufzeit)

	type Wochentag int



	const (
		Dezember float32 = 12.12
		Jannuar  int     = 1
	)

	const DezemberZwei = 12.12

	var montag Wochentag = Montag

	fmt.Printf("Der Montag hat den Wert %d. \n", montag)
	fmt.Printf("Die Konstante Mittwoch hat den Wert %d. \n", Mittwoch)
}
