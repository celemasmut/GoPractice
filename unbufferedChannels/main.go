package main

import (
	"fmt"
	"time"
)

//unbuffered channel means it can take only one value

func main() {
	myChannel := make(chan int) // name - channel - datatype

	//send in a gorutine
	go func() {
		myChannel <- 2 //giving an int value to my channel
	}()

	//sniff

	value := <-myChannel //taking the number value from the channel

	fmt.Println(value)

	//send in a gorutine
	go func() {
		myChannel <- 1 //giving an int value to my channel
	}()
	time.Sleep(time.Second * 2)

	value = <-myChannel //taking the number value from the channel

	fmt.Println(value)

	fmt.Println(myChannel)
}
