package main

import "fmt"

func main() {
	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// switch sin condiciones
	valor := 2
	switch {
	case valor > 100:
		fmt.Println("Es mayor a 100")
	case valor < 0:
		fmt.Println("Es menor que 0")
	default:
		fmt.Println("No cumple ninguna")
	}
}
