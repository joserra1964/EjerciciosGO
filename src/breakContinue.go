package main

import "fmt"

func main() {
	// Defer
	defer fmt.Println("Hola diferido")
	fmt.Println("Mundo")

	// Continue y break
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", i)

		// Continue
		if i == 2 {
			fmt.Println("Es valor 2 ")
			continue
		}

		// Break
		if i == 7 {
			fmt.Println("Es Break")
			break
		}
	}

}
