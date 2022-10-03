/* Ejercicio 1 - Registro de estudiantes
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el
detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y
que tenga un método detalle
*/

package main

import "fmt"

type Alumnos struct {
	name     string
	surname  string
	document int
	date     string
}

func (a Alumnos) detail() {
	fmt.Println("Nombre:", a.name)
	fmt.Println("Apeelido:", a.surname)
	fmt.Println("Documento:", a.document)
	fmt.Println("Fecha de ingreso:", a.date)

}

func main() {

	alumno1 := Alumnos{
		name:     "Juan",
		surname:  "Pedetti",
		document: 4698562,
		date:     "19/09/2022",
	}

	alumno2 := Alumnos{"Martin", "Guglielmone", 490498, "19/09/2022"}

	alumno1.detail()
	alumno2.detail()

}
