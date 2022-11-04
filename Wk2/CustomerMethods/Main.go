package main

import "fmt"

func main() {
	customer1 := &customer{
		Fname:    "Michael",
		Lname:    "Jordan",
		Username: "MJ2020",
		Password: "1234567",
		Email:    "MJ2020@gmail.com",
		Phone:    12345678,
		Address:  "18227 Capstan Greens Road Cornelius, NC 28031.",
	}
	AllUserInformation(customer1)
	fmt.Println(customer1.userCredentials())
	fmt.Println(customer1.userAddress())
}
