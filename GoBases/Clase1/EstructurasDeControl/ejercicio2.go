package main

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

	} else if  edad <= 22 {

		fmt.Println("Prestamo NO otorgado: debe ser mayor a 22 anos")


	} 




}