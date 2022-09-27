/* Ejercicio 2 - Leer archivo
La misma empresa necesita leer el archivo almacenado,
para ello requiere que: se imprima por pantalla mostrando los valores
 tabulados, con un título (tabulado a la izquierda para el ID y a la derecha
	para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe
	visualizar el total (Sumando precio por cantidad)

Ejemplo:

ID                            Precio  Cantidad
111223                      30012.00         1
444321                    1000000.00         4
434321                         50.50         1
                          4030062.50
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	filePath := "./stock.csv"

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	dataS := string(data)

	lines := strings.Split(dataS, "\n")

	for _, line := range lines {
		if len(line) > 0 {
			dataLine := strings.Split(line, ";")
			fmt.Printf("%s\t\t%s\t%s\n", dataLine[0], dataLine[1], dataLine[2])

		}

	}

}
