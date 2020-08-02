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

const password = "admin_1234"
var yourPassword string

func main() {
	for i := 0; ; i+=1 {
		fmt.Println("Type your password!")
		fmt.Scan(&yourPassword)
		if yourPassword == password {
			fmt.Println("Your password is correct! Welcome to the System")
			break
		}else {
			fmt.Println("Type your Password again!")
		}

	}

}
