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
	fmt.Println("END") //heavy and this text will start at the same time so END will be shown first
	//time.Sleep(time.Second * 5) //as I have this time the f.heavy() works
	select {} // with this the f.heavy will run not stopping

}
