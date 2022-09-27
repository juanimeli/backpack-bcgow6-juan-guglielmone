/*
Ejercicio 2 - Calcular promedio
Un colegio necesita calcular el promedio (por alumno)
de sus calificaciones. Se solicita generar una función en la cual se
le pueda pasar N cantidad de enteros y devuelva el promedio y un error
en caso que uno de los números ingresados sea negativo
*/

package main

import (
	"errors"
	"fmt"
)

func avrg(notas ...float64) (prom float64, err error) {

	total := 0.0

	for _, valor := range notas {
		if valor < 0 {
			err = errors.New("Las notas deben ser mayor a cero.")
		}
		total += valor
	}
	prom = total / float64(len(notas))
	return

}

func main() {

	res, err := avrg(20.0, 30.5, 51, 39)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
