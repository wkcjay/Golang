package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Println("Please enter a positive number: ")
	fmt.Scanln(&x)
	fmt.Println("Factorial is: ", factorial(x))
}
func factorial(n int) uint64 {
	var factVal uint64 = 1
	if n < 0 {
		fmt.Println("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= uint64(i)
		}
	}
	return factVal
}
