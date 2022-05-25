package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Jose"] = 14
	m["Pepito"] = 20
	m["Juanita"] = 50
	fmt.Println(m)

	// Recorrer un map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar valor
	valor, ok := m["Jose"]
	fmt.Println(valor, ok)
	valor2, ok2 := m["Joser"]
	fmt.Println(valor2, ok2)

}
