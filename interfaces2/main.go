package main

import "fmt"

//Persona es una estructura
type Persona struct {
	nombre string
	email  string
	edad   int
}

//Moderador es una estructura
type Moderador struct {
	Persona
	Foro string
}

//AbrirForo es una funcion que implementa Moderador
func (m Moderador) AbrirForo() {
	fmt.Println("Abrir Foro")
}

//CerrarForo es una funcion que implementa Moderador
func (m Moderador) CerrarForo() {
	fmt.Println("Cerrar foro")
}

//Administrador es una estructura
type Administrador struct {
	Persona
	Seccion string
}

//CrearForo es una funcion que implementa Administrador
func (a Administrador) CrearForo() {
	fmt.Println("Crear Foro")
}

//EliminarForo es una funcion que implementa Administrador
func (a Administrador) EliminarForo() {
	fmt.Println("Eliminar foro")
}

//Nombre es una funcion de interfaz Usuario que implemente persona
func (p Persona) Nombre() string {
	return p.nombre
}

//Email es una funcion de interfaz Usuario que implemente persona
func (p Persona) Email() string {
	return p.email
}

//Usuario es una interfaz con dos funciones
type Usuario interface {
	Nombre() string
	Email() string
}

//Presentarse es una funcion que recibe una persona
func Presentarse(u Usuario) {
	fmt.Println("Nombre: ", u.Nombre())
	fmt.Println("Email: ", u.Email())
}

func main() {
	Celeste := Persona{"Celeste", "12334@mail.com", 25}
	Presentarse(Celeste)
	Lucho := Persona{"Luciano", "5354343@mail.com", 28}
	Luciano := Administrador{Lucho, "Psicologia"}
	Presentarse(Luciano)

	var user Usuario
	user = Luciano
	fmt.Println("i: ", user)
	fmt.Println("i: ", user.Email())
	fmt.Println("i: ", user.Nombre())

}
