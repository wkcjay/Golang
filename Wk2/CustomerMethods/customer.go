package main

import "fmt"

type customer struct {
	Fname    string
	Lname    string
	Username string
	Password string
	Email    string
	Phone    int
	Address  string
}

func (c *customer) userCredentials() (string, string) {
	return c.Username, c.Password
}
func (c *customer) userAddress() string {
	return c.Address
}
func AllUserInformation(customer *customer) {
	fmt.Println(customer)
}
