package main

import (
	"fmt"
	"sync"
)

func main() {

	//wait groups
	wg := &sync.WaitGroup{}
	//add done and wait functionality
	wg.Add(2)

	//here we have an anonymous function. it does no contain any name
	//useful to create an inline function
	go func() {
		fmt.Println("hello")
		wg.Done()
	}()
	go func() {
		fmt.Println("world")
		wg.Done()
	}() //as I have both f() with go. none of them will show
	wg.Wait()
	fmt.Println("END")
	//select {} //this will cause an indefinitly run ending with deadlock
}
