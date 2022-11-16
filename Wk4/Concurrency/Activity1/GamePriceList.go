package main

import (
	"fmt"
	"sync"
)

type game struct {
	gamelist map[string]float64
	m        sync.Mutex
}

func (g *game) add() {
	var name string
	var price float64
	fmt.Println("Please enter the game name: ")
	fmt.Scanln(&name)
	fmt.Println("Please enter the price: ")
	fmt.Scanln(&price)
	_, exists := g.gamelist[name]
	if exists {
		fmt.Println(name + " has already registered")
	} else {
		// g.m.Lock()
		g.gamelist[name] = price
		// g.m.Unlock()
		fmt.Println("The game item is successfully added!")
	}
}

func (g *game) display() {
	// g.m.Lock()
	fmt.Println("Game Library: ")
	for k, v := range g.gamelist {
		fmt.Println(k, v)
	}
	// g.m.Unlock()
}

func manageGameItems(g *game, wg *sync.WaitGroup) {
	defer wg.Done()
	g.m.Lock()
	var option int
	fmt.Printf("Select an option:\n1. Display all games\n2. Add new game\n-1. Exit\n")
	fmt.Scanln(&option)
	switch option {
	case 1:
		g.display()
	case 2:
		g.add()
	}
	g.m.Unlock()
}
func main() {

	g := &game{gamelist: map[string]float64{}}
	var wg sync.WaitGroup
	wg.Add(2)
	//user 1
	go manageGameItems(g, &wg)
	//user 2
	go manageGameItems(g, &wg)
	wg.Wait()
	fmt.Println()
}
