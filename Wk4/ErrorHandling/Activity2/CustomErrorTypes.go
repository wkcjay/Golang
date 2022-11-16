package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Println("Please enter the first length of the triangle: ")
	fmt.Scanln(&a)
	fmt.Println("Please enter the second length of the triangle: ")
	fmt.Scanln(&b)
	fmt.Println("Please enter the third length of the triangle: ")
	fmt.Scanln(&c)
	area, err := createTriangle(a, b, c)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(area)
	}
}

var invalidTriangleError = errors.New("Invalid Triangle")
var invalidSideError = errors.New("Invalid Side")

func createTriangle(a float64, b float64, c float64) (float64, error) {
	var area float64
	if a < 0 || b < 0 || c < 0 {
		return area, invalidSideError
	} else if a+b <= c || b+c <= a || c+a <= b {
		return area, invalidTriangleError
	} else {
		s := (a + b + c) / 2
		area = math.Sqrt(s * (s - a) * (s - b) * (s - c))
		return area, nil
	}
}
