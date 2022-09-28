/*
Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados
basándose en la cantidad de horas trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados
por mes y la categoría, y que devuelva su salario.
*/

package main

import (
	"errors"
	"fmt"
)

const (
	unC    = 1000
	unB    = 1500
	unA    = 3000
	pctgeB = 20
	pctgeA = 50
)

func salario(minMensual int, categ string) (sueldo int, err error) {

	if minMensual < 0 {
		err = errors.New("Debe ingresar minutos mayor a 0")
	}

	hrs := minMensual / 60
	var sueldoBase int

	switch categ {
	case "C":
		sueldo = unC * hrs
	case "B":
		sueldoBase = unB * hrs
		sueldo = sueldoBase + sueldoBase*pctgeB/100
	case "A":
		sueldoBase = unA * hrs
		sueldo = sueldoBase + sueldoBase*pctgeA/100
	default:
		err = errors.New("Categoria invalida. Valores validos: A, B o C")
	}

	return

}

func main() {

	sueldo, err := salario(65, "C")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sueldo)
	}

	sueldo1, err1 := salario(700, "B")

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(sueldo1)
	}

	sueldo2, err2 := salario(400, "A")

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(sueldo2)
	}

	sueldo3, err3 := salario(-400, "A")

	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(sueldo3)
	}

	sueldo4, err4 := salario(400, "D")

	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println(sueldo4)
	}

}
