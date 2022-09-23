package main

import "fmt"

var temperatura float32
var humedad int
var presion float32

func main () {

	temperatura = 20.4
	humedad = 52
	presion = 1021.1

	fmt.Println(temperatura , "C")
	fmt.Println(humedad , "%")
	fmt.Println(presion , "hPa")
}