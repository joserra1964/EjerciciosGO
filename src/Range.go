package main

import (
	"fmt"
	"strings"
)

func esPalindromo(text string) {
	var textReverse string
	textMin := strings.ToLower(text)
	for i := len(textMin) - 1; i >= 0; i-- {
		textReverse += string(textMin[i])
	}

	if textMin == textReverse {
		fmt.Println(text, ": es palindromo")
	} else {
		fmt.Println(text, ": no es palindromo")
	}

}

func main() {
	slice := []string{"hola", "que", "hace"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	// Palíndromos : ama, amor a roma
	esPalindromo("amA")
	esPalindromo("Asa")
	esPalindromo("casa")
	esPalindromo("arma")

}
