package main

import "fmt"

func main() {
	base := 50
	altura := 10
	areaTriangulo := base * altura / 2
	fmt.Println("Area del triángulo:", areaTriangulo)
	AristaSuperior := 30
	areaTrapecio := base*altura - ((base - AristaSuperior) * altura / 2)
	fmt.Println("Area Trapecio: ", areaTrapecio)

}
