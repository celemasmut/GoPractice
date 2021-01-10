package main

import (
	"fmt"
	"time"
)

func heavy() {
	time.Sleep(time.Second * 5)
}

func main() {
	//go es a word to make work the concurrency
	go heavy()
	fmt.Println("END")

}
