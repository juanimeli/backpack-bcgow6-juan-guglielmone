package main

/*
Ejercicio 1 - Impuestos de salario #1
1.En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
2.Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el
salario ingresado no alcanza el mínimo imponible" y y lánzalo en caso de que “salary” sea menor
a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.

Ejercicio 2 - Impuestos de salario #2

Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de
 “Error()”,  se implemente “errors.New()”.

 Ejercicio 3 - Impuestos de salario #3
Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error
reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible (el mensaje
mostrado por consola deberá decir: “error: el mínimo imponible es de 150.000 y el salario ingresado
es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).

*/

import (
	"errors"
	"fmt"
	"os"
)

type salaryError struct {
	msg string
}

func (err *salaryError) Error() string {
	return fmt.Sprintf("%v", err.msg)
}

func salaryErrorTest(salary int) (res string, err error) {
	if salary < 150000 {
		err = &salaryError{"error: el salario ingresado no alcanza el minimo imponible"}
		return
	}
	res = "Ej1: Debe pagar impuesto"
	return
}

func salaryErrorTest2(salary int) (res string, err error) {
	if salary < 150000 {
		err = errors.New("error: el salario ingresado no alcanza el minimo imponible")
		return
	}
	res = "Ej2: Debe pagar impuesto"
	return
}

func salaryErrorTest3(salary int) (res string, err error) {
	if salary < 150000 {
		err = fmt.Errorf("error: el salario ingresado no alcanza el minimo imponible y el salario ingresado es %d", salary)
		return
	}
	res = "Ej3: Debe pagar impuesto"
	return
}
func main() {

	salary := 200_000

	result, err := salaryErrorTest(salary)

	if err != nil {

		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)

	salary2 := 200_000
	result2, err2 := salaryErrorTest2(salary2)

	if err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}

	fmt.Println(result2)

	salary3 := 200_000
	result3, err3 := salaryErrorTest3(salary3)

	if err3 != nil {
		fmt.Println(err3)
		os.Exit(1)
	}

	fmt.Println(result3)

}
