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

import "fmt"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minFun(calificaciones ...float64) (minC float64) {

	for i, calif := range calificaciones {

		if i == 0 {
			minC = calif
		} else if calif < minC {
			minC = calif
		}
	}
	return
}

func maxFun(calificaciones ...float64) (maxC float64) {
	for i, calif := range calificaciones {

		if i == 0 {
			maxC = calif
		} else if calif > maxC {
			maxC = calif
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

//func handler(stadistic string) func( notas ...float64) float64 {
//	switch stadistic {
//	case minimum:
//		return minFun
//case average:
//	return promFun
//case maximum:
//	return maxFun
//default:

//	}
//}

func main() {

	fmt.Println(minFun(23, 45, 20, 100, 112, 56, 56))
	fmt.Println(maxFun(23, 45, 20, 100, 112, 56, 56))
	fmt.Println(promFun(23, 45, 20, 100, 112, 56, 56))

}
