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
