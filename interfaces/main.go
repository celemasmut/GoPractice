package main

import "fmt"

//User es una interfaz
type User interface {
	Permissions() int //1 - 5
	Name() string
}

//Admin es una estructura
type Admin struct {
	name string
}

//Editor es una estructura
type Editor struct {
	name string
}

//Permissions es una funcion de interfaz que la implemente Admin
func (this Admin) Permissions() int {
	return 5
}

//Name es una funcion de interfaz que la implemente Admin
func (this Admin) Name() string {
	return this.name
}

//Permissions es una funcion de interfaz que la implemente Editor
func (this Editor) Permissions() int {
	return 3
}

//Name es una funcion de interfaz que la implemente Admin
func (this Editor) Name() string {
	return this.name
}

//auth es una funcion que generica
func auth(myUser User) string {
	if myUser.Permissions() > 4 {
		return myUser.Name() + " tiene permisos de administrador"
	} else if myUser.Permissions() == 3 {
		return myUser.Name() + " tiene permisos de editor"
	}
	return ""
}

func main() {
	admin := Admin{"Celeste"}
	editor := Editor{"Luciano"}
	//arreglo de usuarios
	users := []User{admin, editor}

	for _, user := range users {
		fmt.Println(auth(user))
	}
}
