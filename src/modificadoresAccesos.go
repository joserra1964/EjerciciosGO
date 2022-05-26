package main

import (
	mipaquete "EjerciciosGO/src/miPaquete"
	"fmt"
)

// Si la primera letra es mayúscula es público y si es minúscula, es de ámbito privado

func main() {
	var miCoche mipaquete.CarPublic
	miCoche.Marca = "Ferrari"
	fmt.Println(miCoche)
}
