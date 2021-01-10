package main

import (
	"fmt"
	"time"
)

func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Hello")
	}
}

func main() {
	//go es a word to make work the concurrency
	go heavy()
	fmt.Println("END")
}
