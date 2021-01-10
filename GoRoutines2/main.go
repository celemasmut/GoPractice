package main

import (
	"fmt"
)

func main() {
	//here we have an anonymous function. it does no contain any name
	//useful to create an inline function
	go func() {
		fmt.Println("hello")
	}()
	go func() {
		fmt.Println("world")
	}() //as I have both f() with go. none of them will show
	fmt.Println("END")
	select {} //this will cause an indefinitly run ending with deadlock
}
