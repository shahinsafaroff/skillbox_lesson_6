package main

import "fmt"

const egg = "яйцо"
var answer string
func main() {
	for i := 0; ; i += 1 {
		fmt.Println("ответ пользователя")
		fmt.Scan(&answer)
		fmt.Println(egg)

		if answer=="надоело" {
			break
		}
	}
	fmt.Println("Победа!!!")

}
