package main

import "fmt"

func main() {
	// Declaración de variables
	hola := "Hola"
	mundo := "mundo"

	// Println
	fmt.Println(hola, mundo)

	// Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d\n", nombre, cursos)
	// si no sabes que tipo de dato ponle v
	fmt.Printf("%v tiene más de %v\n", nombre, cursos)

	// Sprintf  genera un string pero no lo muestra por consola
	mensaje := fmt.Sprintf("%s tiene más de %d", nombre, cursos)
	fmt.Println(mensaje)

	// Tipo datos
	fmt.Printf("hola: %T\n", hola)
	fmt.Printf("cursos: %T\n", cursos)
	fmt.Printf("mensajeL: %T\n", mensaje)
}
