package main

import "fmt"

func main() {
	var randomNumber, summ int
	for i := 0; i < 4; i += 1 {
		fmt.Println("Type the random number")
		fmt.Scan(&randomNumber)
		summ += randomNumber
		}
	fmt.Println(summ)
}
