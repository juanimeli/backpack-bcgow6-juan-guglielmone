package main

/*
Ejercicio 3 - A qué mes corresponde
Realizar una aplicación que contenga una variable con el número del mes.
Según el número, imprimir el mes que corresponda en texto.
¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
Ej: 7, Julio
*/

import "fmt"

func main() {

	month := 10

	switch month {

	case 1:
		fmt.Println(month, "Enero")
	case 2:
		fmt.Println(month, "Febrero")
	case 3:
		fmt.Println(month, "Marzo")
	case 4:
		fmt.Println(month, "Abril")
	case 5:
		fmt.Println(month, "Mayo")
	case 6:
		fmt.Println(month, "Junio")
	case 7:
		fmt.Println(month, "Julio")
	case 8:
		fmt.Println(month, "Agosto")
	case 9:
		fmt.Println(month, "Septiembre")
	case 10:
		fmt.Println(month, "Octubre")
	case 11:
		fmt.Println(month, "Noviembre")
	case 12:
		fmt.Println(month, "Diciembre")
	default:
		fmt.Println("Desconocido")
	}
}
