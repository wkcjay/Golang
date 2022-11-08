package main

import (
	"fmt"
)

var customer map[string]string

func main() {
	//Creating new hash table for customer
	customer = make(map[string]string)
	var x int
	var Email string
	var inputpw string
	fmt.Println("1.Sign in\n2.Sign Up\n3.Delete Account")
	fmt.Scanln(&x)
	fmt.Println("Please enter your email: ")
	fmt.Scanln(&Email)
	fmt.Println("Please enter your password: ")
	fmt.Scanln(&inputpw)
	switch {
	case x == 1:
		pw, err := get(Email)
		if err == false {
			fmt.Println("Email does not exist")
		} else if inputpw != pw {
			fmt.Println("You have entered the wrong pw")
		} else {
			fmt.Println("Welcome")
		}
	case x == 2:
		_, err := get(Email)
		if err == true {
			fmt.Println("Email existed")
		} else {
			put(Email, inputpw)
			fmt.Println("Account successfully created")
		}

	case x == 3:
		pw, err := get(Email)
		if err == false {
			fmt.Println("Email does not exist")
		} else if inputpw != pw {
			fmt.Println("You have entered the wrong pw")
		} else {
			//delete particular key-value from the hash table
			delete(customer, Email)
			fmt.Println("Account Deleted!")
		}
	}
}

// search key to get value inside hash table
func get(Email string) (string, bool) {
	pw, err := customer[Email]
	return pw, err
}

// insert new key-value into the hash table
func put(Email string, pw string) {
	customer[Email] = pw
}
