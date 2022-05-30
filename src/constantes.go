package main

import "fmt"

const a string = "hola Mundo"
const b = 2.14
const c int = 4

const (
	PI     = 3.14
	cadena = "una Cadena"
	A      = 1
)

func main() {

	fmt.Println(PI)
	fmt.Println(cadena)
	fmt.Println(A)
	fmt.Printf("%T\n", A)

	fmt.Println(a, b, c)

}
