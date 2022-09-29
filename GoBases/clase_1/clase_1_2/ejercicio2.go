package main

/*
Ejercicio 2 - Préstamo
Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga
préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de
un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés
a los que su sueldo es mejor a $100.000.
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.

Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/

import "fmt"

func main() {

	edad := 24
	situacion := "empleado"
	antiguedad := 2
	sueldo := 120000

	if edad > 22 {

		if situacion == "empleado" {
			if antiguedad > 1 {

				if sueldo > 100000 {

					fmt.Println("Prestamo sin interes otorgado")

				} else if sueldo <= 100000 {

					fmt.Println("Prestamo otorgado con tasa de interes")

				}
			} else if antiguedad <= 1 {

				fmt.Println("Prestamo NO otorgado: su antiguedad debe ser mayor a 1 ano")

			}
		} else {

			fmt.Println("Prestamo NO otorgado: debe contar con un empleo")
		}

	} else if edad <= 22 {

		fmt.Println("Prestamo NO otorgado: debe ser mayor a 22 anos")

	}

}
