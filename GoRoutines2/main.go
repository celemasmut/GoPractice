package main

import (
	"fmt"
)

func main() {
	//here we have an anonymous function. it does no contain any name
	//useful to create an inline function
	func() {
		fmt.Println("hello")
	}()
	fmt.Println("END")

}
