package main

import "fmt"

func main() {
	var vtipo1 int = 1
	var vtipo2 = 2
	vtipo3 := 3
	fmt.Println(vtipo1, vtipo2, vtipo3)
	var s1 string = "Una cadena larga"
	var s2 string = "Otra cadena"
	var i1 int = 1
	var b1 bool = false
	fmt.Println(s1, s2, i1, b1)
	fmt.Printf("\n%T \t %T \t %T \t %T \n", s1, s2, i1, b1)

	var a, b = 6, "Hola"
	c, d := 7, "Mundo"
	fmt.Println(a, b, c, d)

	var (
		e int
		f int    = 1
		g string = "Hola"
		h *int   = &e
	)
	fmt.Println(e, f, g, h)
	fmt.Printf("\n%T \t %T \t %T \t %T \n", e, f, g, h)
}
