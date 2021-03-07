package model

import "fmt"

//CreateTODO is the function thar creat the table todo
func CreateTODO() error {
	insertQ, err := con.Query("INSERT INTO TODO VALUES (?, ?)", "Celes", "This api")
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
