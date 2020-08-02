package main

import (
	"fmt"
	"math"
)

var randomNumber, pow, exponential float64
func main() {
	for {
		fmt.Println("Type the random number: ")
		fmt.Scan(&randomNumber)
		fmt.Println("Type the exponential: ")
		fmt.Scan(&pow)
		if pow > 0 {
			exponential = math.Pow(randomNumber, pow)
			fmt.Println(exponential)
			return
		} else {
			fmt.Println("Your exponento is below or equal zero")
		}
	}
}