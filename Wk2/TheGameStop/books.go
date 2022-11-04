package main

import "fmt"

type book struct {
	title string
	price int
}

func (b *book) Print() {
	fmt.Println(fmt.Sprintf("%s, $%d", b.title, b.price))
}
