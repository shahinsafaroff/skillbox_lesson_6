package main

import "fmt"

func main() {
	var word string
	fmt.Scan(&word)
	for i := 0; i < 3; i += 1 {
		fmt.Println(word)
	}
}