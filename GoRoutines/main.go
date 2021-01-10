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

func superHeavy() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("Super heavy")
	}
}

func main() {
	//go es a word to make work the concurrency
	go heavy()
	superHeavy() // as this has not ending the next line will never be shown
	fmt.Println("END")
	//time.Sleep(time.Second * 5) //as I have this time the f.heavy() works
	select {} // with this the f.heavy will run not stopping

}
