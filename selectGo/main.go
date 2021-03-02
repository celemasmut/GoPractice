package main

import (
	"fmt"
	"os"
	"time"
)

//Select is a function
func Select(myChannel chan int, quits chan struct{}) {

	//similar to switch case syntax
	//the difference is that select is for channel asyncronic
	for {
		time.Sleep(time.Second)
		select {
		case <-myChannel:
			fmt.Println("received")
		case <-quits:
			fmt.Println("Quit")
			os.Exit(0) // we can not use break. Use this instead.
		}
	}
}

func main() {
	myChannel := make(chan int)
	quits := make(chan struct{})
	go func() {
		myChannel <- 1
		quits <- struct{}{}
	}()
	go Select(myChannel, quits)

	//up to this we have nothing
	select {} // with this almost works but we have fatal error-deadlock

}
