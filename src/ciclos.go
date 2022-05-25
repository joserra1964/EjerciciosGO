package main

import "fmt"

func main() {

	// For condicional
	for i := 0; i < 10; i++ {
		fmt.Printf("%d - \t", i)
	}
	fmt.Println()

	// For while
	contador := 0
	for contador < 10 {
		fmt.Printf("%d\t", contador)
		contador++
	}
	fmt.Println()

	// For forever
	siempre := 0
	for {
		fmt.Printf("%d\t", siempre)
		siempre++
		// *if siempre > 20 {
		//	break
		// }
		// ahora se detiene con crtl + C
		fmt.Println()
	}

}
