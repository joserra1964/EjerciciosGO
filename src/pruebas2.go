package main

import "fmt"

func main() {
	fmt.Printf("\n El valor es %v, %b, %d, %o, %x", 255, 255, 255, 255, 255)
	fmt.Println()
	var s string = "abcdefgh"
	fmt.Printf("\n El valor es %v, %s, %q, %x, %X", s, s, s, s, s)
	fmt.Println()
	s = "ABCDEFGH"
	fmt.Printf("\n El valor es %v, %s, %q, %x, %X", s, s, s, s, s)
	fmt.Println()
	var ps = &s
	fmt.Printf("\n El valor es %v, %p", ps, ps)
	fmt.Println()
	var ps2 *string
	ps2 = ps
	fmt.Printf("\n El valor es %v, %p", ps2, ps2)
	fmt.Println()
	var ps3 string
	ps3 = *ps
	fmt.Printf("\n El valor es %v, %p", ps3, ps3)
	fmt.Println()
}
