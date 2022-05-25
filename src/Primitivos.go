package main

import (
	"fmt"
	"math"
)

func main() {
	// Numeros enteros
	fmt.Println("Numeros enteros")
	// int = depende del OS (32 o 64 bits)
	fmt.Println("int = depende del OS (32 o 64 bits)")
	// int8 = 8 bits = -128 a 127
	fmt.Println("int8 = 8 bits = ", -math.Exp2(7), "a ", math.Exp2(7)-1)
	// int16 = 16 bits = -2^15 a 2^15 -1
	fmt.Println("int16 = 16 bits = ", -math.Exp2(15), "a ", math.Exp2(15)-1)
	// int32 = 32 bits = -2^31 a 2^31 -1
	fmt.Println("int32 = 32 bits = ", -math.Exp2(31), "a ", math.Exp2(31)-1)
	// int64 = 64 bits = -2^63 a 2^63 -1
	fmt.Println("int64 = 64 bits = ", -math.Exp2(63), "a ", math.Exp2(63)-1)

	// Numeros enteros positivos
	fmt.Println("Números enteros positivos")
	// uint = depende del OS (32 o 64 bits)
	fmt.Println("depende del OS (32 o 64 bits)")
	// uint8 = 8 bits = 0 a 2^8 -1
	fmt.Println("uint8 = 8 bits = ", 0, "a ", math.Exp2(8)-1)
	// uint16 = 16 bits = 0 a 2^16 -1
	fmt.Println("uint16 = 16 bits = ", 0, "a ", math.Exp2(16)-1)
	// uint32 = 32 bits = 0 a 2^32 -1
	fmt.Println("uint32 = 32 bits = ", 0, "a ", math.Exp2(32)-1)
	// uint64 = 64 bits = 0 a 2^64 -1
	fmt.Println("uint64 = 64 bits = ", 0, "a ", math.Exp2(64)-1)

	// Numeros decimales
	fmt.Println("Numeros decimales")
	// float32 = 32 bits = +/- 1.18 10^(-38) a +/- 3.4 10^(38)
	fmt.Println("float32 = 32 bits = +/-", 1.18e-38, "a  +/-", 3.4e38)
	// float64 = 64 bits = +/- 2.23 10^(-308) a +/- 1.8 10^(308)
	fmt.Println("float64 = 64 bits = +/-", 2.23e-308, "a  +/-", 1.79e308)

	// Textos y boleanos
	fmt.Println("Textos y boleanos")
	// string =""
	fmt.Println("string =")
	// bool = true o false
	fmt.Println("bool = true o false")

	// Números complejos
	fmt.Println("Números complejos")
	// Complex64 = Real e imaginario float32
	fmt.Println("// Complex64 = Real e imaginario float32")
	// Complex128 = Real e imaginario float64
	fmt.Println("Complex128 = Real e imaginario float64")
	// Ejemplo c:= 10 + 8i

}
