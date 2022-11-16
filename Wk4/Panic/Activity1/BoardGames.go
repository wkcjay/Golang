package main

import (
	"errors"
	"fmt"
)

var boardgames = []string{
	"Carcasssone", "Wildlife Safari", "Civilization",
}

func main() {
	fmt.Println("The boardgames are :")
	for i := 0; i < len(boardgames); i++ {
		fmt.Println(i+1, ":", boardgames[i])
	}
	var x int
	fmt.Println("What is your favourite game?")
	fmt.Scanln(&x)
	print(x)
}

func print(x int) {
	defer func() {
		if err := recover(); err != nil {
			err = errors.New("You have entered an invalid choice. Value should be between 1 and 3")
			fmt.Println(err)
		}
	}()
	fmt.Println("Oh I see. So your favourite game is:", boardgames[x-1])
}
