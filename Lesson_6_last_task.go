package main

import "fmt"

func main() {
	summ := 0

	for digit := 0; digit <= 40; digit += 1 {
		fmt.Println("Divided to four", digit)

		if digit % 4 != 0 {
			continue
		}
		summ += digit
		fmt.Println(summ)
	}

}
