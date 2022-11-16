package main

import (
	"fmt"
	"math"
)

func sum(input int, pow float64, ch chan int) {
	var value int
	cal := 0
	for input > 0 {
		value = input % 10
		cal += int(math.Pow(float64(value), pow))
		input /= 10
	}
	ch <- cal
}
func main() {
	var input int
	sqch := make(chan int)
	cubch := make(chan int)
	fmt.Println("Please key in a number: ")
	fmt.Scanln(&input)
	go sum(input, 2, sqch)
	go sum(input, 3, cubch)
	squares, cubes := <-sqch, <-cubch
	fmt.Println("squares = ", squares)
	fmt.Println("cubes = ", cubes)
	fmt.Println("output = squares + cubes = ", squares+cubes)
}
