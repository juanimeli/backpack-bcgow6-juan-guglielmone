package main

/*
Ejercicio 1 - Test Unitario Restar
Para el método Restar() visto en la clase, realizar el test unitario correspondiente. Para esto:
1. Dentro de la carpeta go-testing crear un archivo calculadora.go con la función a probar.
2. Dentro de la carpeta go-testing crear un archivo calculadora_test.go con el test diseñado.

Ejercicio 2 - Test Unitario Método Ordenar
Diseñar un método que reciba un slice de enteros y los ordene de forma ascendente, posteriormente
diseñar un test unitario que valide el funcionamiento del mismo.
1. Dentro de la carpeta go-testing crear un archivo ordenamiento.go con la función a probar.
2. Dentro de la carpeta go-testing crear un archivo ordenamiento_test.go con el test diseñado.

Ejercicio 3 - Test Unitario Método Dividir
Para el Método Dividir, visto en la clase:

Cambiar el método para que no sólo retorne un entero sino también un error. Incorporar una
validación en la que si el denominador es igual a 0,  retorna un error cuyo mensaje sea
“El denominador no puede ser 0”. Diseñar un test unitario que valide el error cuando se
invoca con 0 en el denominador.
1. Dentro de la carpeta go-testing crear un archivo dividir.go con la función a probar.
2. Dentro de la carpeta go-testing crear un archivo dividir test.go con el test diseñado.
*/

import (
	"fmt"

	"github.com/github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goTesting/clase_1_2/dividir"
	"github.com/github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goTesting/clase_1_2/ordenamiento"
)

func main() {

	s := []int{3, 4, 1, 6, 8, 2}

	ordenamiento.OrderSlice(s)

	fmt.Println(s)

	num1 := 10
	den := 2

	res, _ := dividir.Dividir(num1, den)
	fmt.Println(res)

}
