package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	wg   sync.WaitGroup
	mine = &bankBalance{amount: 100.00, currency: "SGD"}
)

type bankBalance struct {
	amount   float64
	currency string
	mu       sync.Mutex
}

func (b *bankBalance) Deposit() {

	b.mu.Lock()
	{
		var amt float64
		fmt.Println("Amount of " + b.currency + " to deposit: ")
		fmt.Scanln(&amt)
		b.amount += amt
		fmt.Printf("Bank Balance: %s %.2f\n", b.currency, b.amount)
	}
	b.mu.Unlock()
}

func (b *bankBalance) Withdraw() {

	b.mu.Lock()
	{
		var amt float64
		fmt.Println("Amount of " + b.currency + " to withdraw: ")
		fmt.Scanln(&amt)
		b.amount -= amt
		fmt.Printf("Bank Balance: %s %.2f\n", b.currency, b.amount)
	}
	b.mu.Unlock()
}

func (b *bankBalance) Display() {
	b.mu.Lock()
	{
		fmt.Printf("Bank Balance: %s %.2f\n", b.currency, b.amount)
	}
	b.mu.Unlock()
}

func processBankBalance(b *bankBalance) {
	defer wg.Done()
	var option int
	fmt.Printf("Select Option:\n1. Deposit\n2. Withdraw\n3. Display Information\n")
	fmt.Scanln(&option)
	if option == 1 {
		b.Deposit()
	} else if option == 2 {
		b.Withdraw()
	} else if option == 3 {
		b.Display()
	} else {
		fmt.Println("Invalid Input")
	}
}
func main() {
	runtime.GOMAXPROCS(2)
	b := mine
	wg.Add(2)
	go processBankBalance(b)
	go processBankBalance(b)
	wg.Wait()
	fmt.Println()
}
