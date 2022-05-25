package main

import "fmt"

func normalFunction(mensaje string) {
	fmt.Println("Hola mundo")
}

func tripleArgumentos(a, b int, c string) {
	fmt.Println(a, b, c)
}

func devolverValor(a int) int {
	return a * 2
}

func devuelveDos(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola Mundo")
	tripleArgumentos(1, 2, "hola")
	valor := devolverValor(2)
	fmt.Println("valor : ", valor)
	valor1, valor2 := devuelveDos(4)
	fmt.Println("valor1", valor1, "   Valor2", valor2)
	_, valor1 = devuelveDos(10)
	fmt.Println("valor: ", valor1)

}
