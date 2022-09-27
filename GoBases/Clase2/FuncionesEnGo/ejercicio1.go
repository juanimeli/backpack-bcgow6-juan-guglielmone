package main

import "fmt"

func impSal(salario float64) float64 {

	if salario > 150000 {
		return salario * 0.27
	} else if salario > 50000 {
		return salario * 0.17
	} else {
		return 0
	}

}

func main() {

	fmt.Println(impSal(30000))
	fmt.Println(impSal(60000))
	fmt.Println(impSal(200000))

}
