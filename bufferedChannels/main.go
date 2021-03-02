package main

import "fmt"

//buffered channels means that it can take more than 1 value concurrently

//Car is an struct
type Car struct {
	Model string
}

func main() {
	myChannel := make(chan Car, 3) // name - channel - datatype - bufferSize

	go func() {

		myChannel <- Car{"one"}
		myChannel <- Car{"two"}
		myChannel <- Car{"three"}
		myChannel <- Car{"four"}

		close(myChannel)

	}()

	for i := range myChannel {
		fmt.Println(i.Model)
	} // if we iterate the channel like this we get a fatal error - deadlock. this is because
	//the channel is not closed. if we close the channel, solved.

}
