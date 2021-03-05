package model

import (
	"database/sql"
	"databse/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@(tcp:localhost:3306")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the DataBase")
	con = db
	return db
}
