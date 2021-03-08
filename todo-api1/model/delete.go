package model

import "fmt"

//DeleteTODO is the function thar erase
func DeleteTODO(name string) error { //receive a name and the name of de db
	insertQ, err := con.Query("DELETE FROM TODO WHERE name=?", name)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
