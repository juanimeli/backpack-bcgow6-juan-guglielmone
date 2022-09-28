package main

/*
Ejercicio 1 - Red social
Una empresa de redes sociales requiere implementar una estructura usuario
con funciones que vayan agregando información a la estructura. Para optimizar y
ahorrar memoria requieren que la estructura de usuarios ocupe el mismo lugar en
memoria para el main del programa y para las funciones.
La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
Y deben implementarse las funciones:
Cambiar nombre: me permite cambiar el nombre y apellido.
Cambiar edad: me permite cambiar la edad.
Cambiar correo: me permite cambiar el correo.
Cambiar contraseña: me permite cambiar la contraseña.
*/

import (
	"fmt"
)

type User struct {
	Name       string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

func (u User) changeName(newName, newSurName string) {
	p1 := &u
	p1.Name = newName
	p1.Apellido = newSurName
}

func (u User) changeAge(newAge int) {
	p1 := &u
	p1.Edad = newAge
}
func (u User) changeEmail(newEmail string) {
	u.Correo = newEmail
}
func (u User) changePassword(newPass string) {
	u.Contraseña = newPass
}

func main() {

	u1 := User{"a", "b", 3, "c", "d"}

	p2 := &u1

	fmt.Println("La direccion de memoria es:", p2)

	fmt.Println(u1.Edad)
	fmt.Println(u1.Name)
	fmt.Println(u1.Apellido)

	u1.changeAge(20)
	u1.changeName("Juan", "Pedetti")

	fmt.Println(u1.Edad)
	fmt.Println(u1.Name)
	fmt.Println(u1.Apellido)

}
