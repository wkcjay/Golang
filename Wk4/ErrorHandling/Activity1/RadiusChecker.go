package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Println("Please enter a radius: ")
	fmt.Scanln(&r)
	area, err := calCircleArea(r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(area)
	}
}

func calCircleArea(r float64) (float64, error) {
	var err error
	var area float64

	if r > 0 {
		area = math.Pi * r * r
	} else {
		err = errors.New("wrong radius input")
	}
	return area, err
}

var invalidTriangleError = errors.New("Invalid Triangle")
var invalidSideError = errors.New("Invalid Side")

func createTriangle(a float64, b float64, c float64) (float64, error) {
	var area float64
	if a < 0 || b < 0 || c < 0 {
		return area, invalidSideError
	} else if a+b < c || b+c < a || c+a < b {
		return area, invalidTriangleError
	} else {
		s := (a + b + c) / 2
		area = math.Sqrt(s * (s - a) * (s - b) * (s - c))
		return area, nil
	}
}
