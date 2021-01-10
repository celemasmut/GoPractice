package main

import (
	"fmt"
	"time"
)

func heavy() {
	time.Sleep(time.Second * 5)
}

func main() {
	heavy()
	fmt.Println("END")

}
