package main

import (
	"fmt"
)

func main() {
	first := getfirst()
	last := getlast()
	fmt.Println("Hello, " + first + last + " !! Nice to meet you")
}

func getfirst() string {
	first := ""
	fmt.Print("Please enter your first name: ")
	fmt.Scanln(&first)
	if first == "" {
		panic("runtime error: first name cannot be nil")
	} else {
		return first
	}
}

func getlast() string {
	last := ""
	fmt.Print("Please enter your first name: ")
	fmt.Scanln(&last)
	if last == "" {
		panic("runtime error: first name cannot be nil")
	} else {
		return last
	}
}
