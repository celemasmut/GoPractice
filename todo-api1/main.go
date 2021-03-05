package main

import (
	"log"
	"net/http"

	"github.com/celemasmut/GoPractice/GoPractice/todo-api1/controller"
	"github.com/celemasmut/GoPractice/GoPractice/todo-api1/model"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	log.Fatal(http.ListenAndServe(":3000", mux))
}

//once we run the main. then we open de localhost on the navegator.
