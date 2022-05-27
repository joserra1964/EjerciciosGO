package main

import (
	"fmt"
	"math"
)

type area2D interface {
	area() float64
}

type cuadrado struct {
	lado float64
}

type rectangulo struct {
	base   float64
	altura float64
}

type circulo struct {
	radio float64
}

func (c cuadrado) area() float64 {
	return c.lado * c.lado
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func (c circulo) area() float64 {
	return 2 * math.Pi * c.radio
}

func calcular(f area2D) {
	fmt.Printf("Tipo %T : ", f)
	fmt.Println("Area: ", f.area())
}

func main() {

	const pi float64 = 3.14
	fmt.Println("El numero pi es : ", pi)
	miCuadrado := cuadrado{lado: 4}
	miRectangulo := rectangulo{base: 10, altura: 5}
	miCirculo := circulo{radio: 10}
	calcular(miCuadrado)
	calcular(miRectangulo)
	calcular(miCirculo)

	// Lista interfaces
	miInterface := []interface{}{"Hola", 12, 4.5, "Adios"}
	fmt.Println(miInterface...)
}
