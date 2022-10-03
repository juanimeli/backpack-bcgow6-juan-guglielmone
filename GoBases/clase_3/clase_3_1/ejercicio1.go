/*
Ejercicio 1 - Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita:
1. Implementar una funcionalidad para guardar un archivo de texto, con la información
 de productos comprados, separados por punto y coma (csv).
2. Debe tener el id del producto, precio y la cantidad.
3. Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/

package main

import (
	"fmt"
	"os"
)

type Producto struct {
	id     int
	price  float64
	amount int
}

func main() {

	prod1 := Producto{
		id:     1,
		price:  15.00,
		amount: 4,
	}
	prod2 := Producto{
		id:     2,
		price:  30.00,
		amount: 6,
	}

	inventario := [2]Producto{prod1, prod2}

	data := "ID;Precio;Cantidad\n"

	for _, prod := range inventario {
		newProd := fmt.Sprintf("%d;%.2f;%d\n", prod.id, prod.price, prod.amount)
		data += newProd
	}

	dataB := []byte(data)

	err := os.WriteFile("./stock.csv", dataB, 0644)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)

}
