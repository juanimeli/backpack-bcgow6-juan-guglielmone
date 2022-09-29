package main

/*
Ejercicio 2 - Clima

Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura
y humedad y presión atmosférica de distintos lugares.
Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura,
humedad y presión de donde te encuentres.
Imprime los valores de las variables en consola.
¿Qué tipo de dato le asignarías a las variables?

*/

import "fmt"

func main() {

	var temperatura float32 = 20.4
	var humedad int = 52
	var presion float32 = 1021.1

	fmt.Println(temperatura, "C")
	fmt.Println(humedad, "%")
	fmt.Println(presion, "hPa")
}
