package main

import "fmt"

func main() {
	word := "Abecedario"
	fmt.Println("Largo de palabra:", len(word))
	for i := 0; i < len(word); i++ {

		fmt.Printf("%c\n",  word[i])
		//fmt.Printf(word[i])
	}

}