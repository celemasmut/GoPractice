package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/celemasmut/GoPractice/GoPractice/todo-api1/controller"
	"github.com/celemasmut/GoPractice/GoPractice/todo-api1/model"
	_ "github.com/go-sql-driver/mysql" //mysql driver
)

func main() {
	mux := controller.Register() // create a new mux
	db := model.Connect()        // connectint to the database
	defer db.Close()             //this to invoke this line at the end
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe(":3000", mux)) // listen and serv on port 3000
}

//once we run the main. then we open de localhost on the navegator.
