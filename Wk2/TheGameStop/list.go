package main

import "fmt"

type Item interface {
	Print()
}
type List []Item

func (l *List) Print() {
	for _, v := range *l {
		fmt.Printf(" %v\n", v)
	}
}
