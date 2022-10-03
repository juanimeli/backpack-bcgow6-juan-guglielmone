package main

import (
	"errors"
	"fmt"
	"os"
)

/* Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
 Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan
  haber muchos más animales que refugiar.

perro necesitan 10 kg de alimento
gato 5 kg
Hamster 250 gramos.
Tarántula 150 gramos.

Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal
especificado y que retorne una función y un mensaje (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

ejemplo:
const (
   dog = "dog"
   cat = "cat"
)
...

animalDog, msg := Animal(dog)
animalCat, msg := Animal(cat)
...
var amount float64
amount+= animaldog(5)
amount+= animalCat(8)
*/

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"

	kgPorPerro     = 10.0
	kgPorGato      = 5.0
	kgPorHamster   = 0.250
	kgPorTarantula = 0.150
)

func kgParaPerro(cantidad int) float64 {
	return kgPorPerro * float64(cantidad)
}

func kgParaGato(cantidad int) float64 {
	return kgPorPerro * float64(cantidad)
}

func kgParaHamster(cantidad int) float64 {
	return kgPorHamster * float64(cantidad)
}

func kgParaTarantula(cantidad int) float64 {
	return kgPorTarantula * float64(cantidad)
}

func Animal(animal string) (func(int) float64, error) {
	switch animal {
	case perro:
		return kgParaPerro, nil
	case gato:
		return kgParaGato, nil
	case hamster:
		return kgParaHamster, nil
	case tarantula:
		return kgParaTarantula, nil
	default:
		msgError := fmt.Sprintf("%s no es un animal registrado en la calculadora", animal)
		return nil, errors.New(msgError)
	}

}

func main() {
	var total float64

	animalesRefugio := [4]int{24, 32, 10, 4}

	alimentoPerros, err := Animal(perro)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	total += alimentoPerros(animalesRefugio[0])

	alimentoGatos, err := Animal(gato)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	total += alimentoGatos(animalesRefugio[1])

	alimentoHamsters, err := Animal(hamster)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	total += alimentoHamsters(animalesRefugio[2])

	alimentoTarantulas, err := Animal(tarantula)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	total += alimentoTarantulas(animalesRefugio[3])

	fmt.Printf("El total de kg para de alimento a comprar para los animales de refugio es: %.2f\n", total)

	_, err1 := Animal("unKnown")
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}
}
