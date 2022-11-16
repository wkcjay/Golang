package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func readfile(name string, ch chan []string) {
	file, err := os.Open(name)
	if err != nil {
		println("Cannot find file specified.")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txt []string

	for scanner.Scan() {
		txt = append(txt, scanner.Text())
	}
	file.Close()
	ch <- txt
	close(ch)
}

func write(c chan []string, name string) {
	var gamelist []string
	ch := make(chan []string)
	go readfile(name, ch)
	time.Sleep(5 * time.Second)

	gamelist = <-ch
	c <- gamelist
	close(c)
}

func print(list []string) {
	for _, s := range list {
		fmt.Println(s)
	}
}

func main() {
	ch := make(chan []string)
	ch2 := make(chan []string)
	ex, err := os.Getwd()
	if err != nil {
		fmt.Println("No directory")
	}
	expath1 := filepath.Dir(ex) + "/gamelist1.txt"
	expath2 := filepath.Dir(ex) + "/gamelist2.txt"
	go write(ch, expath1)
	go write(ch2, expath2)
	timeout := time.After(10 * time.Second)
loop:
	for {
		select {
		case v1 := <-ch:
			print(v1)
		case v2 := <-ch2:
			print(v2)
		case <-timeout:
			break loop
		}
	}
}
