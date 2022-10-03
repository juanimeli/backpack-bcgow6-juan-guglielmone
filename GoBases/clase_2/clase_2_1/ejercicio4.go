/* Ejercicio 4 - Calcular estadísticas
Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones
de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio)
y que devuelva otra función ( y un mensaje en caso que el cálculo no esté definido) que se le puede pasar una
cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior
Ejemplo:

const (
   minimum = "minimum"
   average = "average"
   maximum = "maximum"
)

...

minFunc, err := operation(minimum)
averageFunc, err := operation(average)
maxFunc, err := operation(maximum)

...

minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
*/

package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minFun(calificaciones ...float64) (minC float64) {

	for i, calificacion := range calificaciones {
		if i == 0 {
			minC = calificacion
		}
		if calificacion < minC {
			minC = calificacion
		}
	}
	return
}

func maxFun(calificaciones ...float64) (maxC float64) {

	for i, calificacion := range calificaciones {
		if i == 0 {
			maxC = calificacion
		}
		if calificacion > maxC {
			maxC = calificacion
		}
	}
	return
}

func promFun(calificaciones ...float64) (prom float64) {

	total := 0.0

	for _, calif := range calificaciones {
		total += calif
	}
	prom = total / float64(len(calificaciones))

	return

}

func calculator(stadistic string) (func(...float64) float64, error) {
	switch stadistic {
	case minimum:
		return minFun, nil
	case average:
		return promFun, nil
	case maximum:
		return maxFun, nil
	default:
		msgError := fmt.Sprintf("La funcion ingresada no existe. Usted ingreso %s", stadistic)
		return nil, errors.New(msgError)

	}
}

func main() {

	minCalification, err := calculator(minimum)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("El minimo de las caficaciones es: %.0f\n", minCalification(45, 23, 76, 34, 12, 6, 34, 65, 98, 34))

	maxCalification, err := calculator(maximum)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("El maximo de las caficaciones es: %.0f\n", maxCalification(45, 23, 76, 34, 12, 6, 34, 65, 98, 34))

	promCalification, err := calculator("average")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("El promedio de las caficaciones es: %.2f\n", promCalification(45, 23, 76, 34, 12, 6, 34, 65, 98, 34))

	unKnown, err := calculator("unknown")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("El minimo de las caficaciones es: %.2f\n", unKnown(45, 23, 76, 34, 12, 6, 34, 65, 98, 34))

}
