package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("Edad Benjamin:", employees["Benjamin"])

	i := 0

	for _,element := range employees {

		if element > 21 {
			i += 1
		} else {
			continue
		}
	}

	fmt.Println("Empleados mayores a 21:",i)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)


}