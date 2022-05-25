package main

import "fmt"

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaración de variables enteras
	base := 12          // los : son porque es la primera vez que la declaramos
	var altura int = 14 // var la declara y el = ya no lleva los : para asignarse
	var area int        // se declara la variable pero no se le asigna valor

	fmt.Println(base, altura, area)

	// Zero values
	var a int     // por defecto le asigna 0
	var b float64 // por defecto le asigna 0
	var c string  // por defecto le asigna un string vacio, no es un valor nulo
	var d bool    // por defecto le asigna false
	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Areaa cuadrado : ", areaCuadrado)

}
