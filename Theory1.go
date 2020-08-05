package main

import "fmt"

const rightAnswer = "25"

func main() {
	userAnswer := ""

	for i := 0; i < 3; i += 1 {
		fmt.Println("Type answer for 5x5: ")
		fmt.Scan(&userAnswer)

		if userAnswer == rightAnswer {
			fmt.Println("Right! You've passed the Exam!")
			return
		} else if userAnswer != rightAnswer {
			fmt.Println("Wrong you should try again!")
		}
	}
	fmt.Println("All 3 tries were wasted! Exam is failed!")
}
package main

import "fmt"

func main() {
	i := 0
	j := 0

	for i < 10 {
		i = i + 1
		j = j + 2
		fmt.Println("Hello World", i)


		if j == 8 {
			break
		}
	}
}
