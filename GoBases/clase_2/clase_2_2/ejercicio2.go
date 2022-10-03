/*
Ejercicio 2 - Matrix
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente
una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es
cuadrática y cuál es el valor máximo.
*/

package main

import (
	"errors"
	"fmt"
)

type Matrix struct {
	matrixRow    int
	matrixColumn int
	matrixValues []float64
}

func (m Matrix) Set() (err error) {

	if m.matrixColumn*m.matrixRow != len(m.matrixValues) {
		err = errors.New("la cantidad de valores debe ser igual los espacios disponibles (colxfilas)")
		return
	}
	return nil

}

func (m Matrix) esCuadratica() bool {
	if m.matrixColumn == m.matrixRow {
		return true
	}
	return false
}

func (m Matrix) Print() {

	if len(m.matrixValues) == 0 {
		fmt.Println("There are no values on this matrix")
	}
	var maxValue float64
	for i, value := range m.matrixValues {
		if i == 0 {
			maxValue = m.matrixValues[0]
		}
		if value > maxValue {
			maxValue = value
		}
	}
	fmt.Printf("The max value is: %.0f\n", maxValue)
	for row := 0; row < m.matrixColumn; row++ {
		fmt.Printf("\t%.0f\n", m.matrixValues[row*m.matrixColumn:row*m.matrixColumn+m.matrixColumn])
	}
}

func main() {

	m := Matrix{
		matrixRow:    5,
		matrixColumn: 4,
		matrixValues: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
	}

	m.Set()
	m.Print()
	fmt.Println(m.esCuadratica())

}
