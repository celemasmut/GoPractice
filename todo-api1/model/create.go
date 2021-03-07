package model

import "fmt"

//CreateTODO is the function thar creat the table todo
func CreateTODO(name, todo string) error { //receive a name and the name of de db
	insertQ, err := con.Query("INSERT INTO TODO VALUES (?, ?)", name, todo)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
