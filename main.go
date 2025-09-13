package main

import (
	"basic_login/models"
	"fmt"
)

func main() {

	var userInput int8

	var initial = "Hi! What do you want to do?\n 0) login\n 1) sign up\n 2)exit"
	server := &models.Server{}
	for {
		fmt.Println(initial)
		fmt.Scan(&userInput)
		switch userInput {
		case 0:
			var id, password string
			fmt.Println("please enter your id")
			fmt.Scan(&id)
			fmt.Println("please enter you password")
			fmt.Scan(&password)

			cli := server.GetClient(id)
			if cli != nil && server.VerifyPassword(password, id) {
				fmt.Println("Login successful")
			} else {
				fmt.Println("Login failed")
			}
		case 1:
			fmt.Println("still developing")
		case 2:
			return
		default:
			fmt.Println("Invalid input")
		}
	}
}
