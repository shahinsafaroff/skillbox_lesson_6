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
		exponential = math.Pow(randomNumber, pow)
		fmt.Println(exponential)
		return
	}
}