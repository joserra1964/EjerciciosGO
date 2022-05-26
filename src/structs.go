package main

import "fmt"

type car struct {
	marca string
	año   int
}

func main() {
	miCoche := car{marca: "Ford", año: 2020}
	fmt.Println(miCoche)

	// Otra manera
	var otroCoche car
	otroCoche.marca = "Ferrari"
	fmt.Println(otroCoche)
}
