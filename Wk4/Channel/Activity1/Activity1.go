package main

import (
	"fmt"
	"time"
)

func hello(done chan bool, num int) {
	fmt.Println("Hello world goroutine", num)
	time.Sleep(4 * time.Second)
	done <- true
}

func main() {
	done := make(chan bool)
	for i := 1; i <= 4; i++ {
		go hello(done, i)
		<-done
	}
	// time.Sleep(4 * time.Second)
	fmt.Println("main function")
}
