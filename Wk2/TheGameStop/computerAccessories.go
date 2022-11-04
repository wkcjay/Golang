package main

import "fmt"

type computerAccessories struct {
	title string
	price int
}

func (ca *computerAccessories) Print() {
	fmt.Println(fmt.Sprintf("%s, $%d", ca.title, ca.price))
}
