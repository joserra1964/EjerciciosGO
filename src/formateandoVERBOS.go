package main

import (
	"fmt"
)

func main() {

	// verbs de formatos
	// %v por defecto, %#v formato sintaxis Go, %T tipo de valor, %% imprime %

	var i = 15.5
	var txt = "Hola Mundo"
	fmt.Printf("%v\t %#v\t %v%%\t %T \t \n", i, i, i, i)
	fmt.Printf("%v\t %#v\t %v%%\t %T \t \n", txt, txt, txt, txt)

	/*	formatos de enteros
		%b base 2, %d base 10, %+d base 10 y signo,
		%o base 8, %O base 8 encabezado por 0o
		%x base 16 minúsculas, 	%X base 16 mayúsculas
		%#x base 156 encabezado 0x
		%4d ancho 4 dígitos justificado derecha, %-4d justic. izquierda,
		%04d relleno con 0
	*/
	var j = 15
	fmt.Printf("%b\t %d\t %+d\t %o\t %O\t %x\t %X\t %4d\t %-4d\t %04d\t \n", j, j, j, j, j, j, j, j, j, j)

	/* 	formateo de cadenas strings
	%s texto plano, %q entre doble comillas, %8s ancho 8 justificado derecha, %-8s idem izquierda
	%x valor como hexadecimal byte, % x idem con expacios
	*/

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%30s\n", txt)
	fmt.Printf("%-30s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	/*	formateando booleanos
		%t valor true o false
	*/

	var v, f bool = true, false
	fmt.Printf("%t\t %t \n", v, f)

	/*	formateando float decimales
		%e notación  científica con e como exponente
		%f decimal sin exponente
		%.2f precisión 2 decimales
		%6.2f 6 dígitos, y 2 decimales
		%g depende si necesita o no dígitos o exponente
	*/

	var z float64 = 3.14151692
	fmt.Printf("%e\t %f\t %.2f\t %6.2f\t %g \n", z, z, z, z, z)

}
