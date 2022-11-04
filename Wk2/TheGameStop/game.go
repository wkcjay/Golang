package main

import "fmt"

type game struct {
	title string
	price float64
}

func (g *game) Print() {
	fmt.Println(fmt.Sprintf("%s: %.2f", g.title, g.price))
}
