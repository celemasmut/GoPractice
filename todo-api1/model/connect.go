package model

import (
	"database/sql"
	"fmt"
	"log"
)

//Con is the global variable to connect db
var con *sql.DB //global connection

//Connect create the connection
func Connect() *sql.DB { //returns an instance to a refer the sql DB object
	dsn := "root:r16c95m.jaja@tcp(localhost:3306)/mysql"
	db, err := sql.Open("mysql", dsn) // opens a database specified by its database driver name and a driver-specific data source name
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the DataBase")
	con = db
	return db
}

// to connect mysql comand i should write on the terminal -> mysql -u root -p"password that I wrote on my note"
