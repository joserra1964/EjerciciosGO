package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// Puerta lógica AND
	if valor1 == 1 && valor2 == 3 {
		fmt.Println("Condición AND es verdad")
	}

	// Puerta lógica OR
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Condición OR es verdad")
	}

	// Convertir texto a número
	valor3, err := strconv.Atoi("es un texto")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Valor3: ", valor3)
}
