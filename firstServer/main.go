package main

import (
	"fmt"
	"net/http" //provides HTTP client and server implementations
)

func main() {

	mux := http.NewServeMux() /*NewServeMux allocates and returns a new ServeMux.
	ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request
	against a list of registered patterns and calls the handler for the pattern that most
	closely matches the URL.*/
	mux.HandleFunc("/getgoing", func(w http.ResponseWriter, r *http.Request) {
		/*A ResponseWriter interface is used by an HTTP handler to construct an HTTP response*/
		/*A Request represents an HTTP request received by a server or to be sent by a client.*/
		fmt.Println("request received")
		w.Write([]byte(" \nHello World")) // writes the data to the connection as part of an HTTP reply
	}) //registers the handler function for the given pattern.
	http.ListenAndServe("localhost:3000", mux) //listenAndServe starts an HTTP server with a given address and handler.
}

// in orden to see first we run the terminal with - go run main.go
//then we use git and run the comand - curl localhost:3000/getgoing . see what happend
