package main

import (
	pk "EjerciciosGO/src/miPaquete"
	"fmt"
)

// Si la primera letra es mayúscula es público y si es minúscula, es de ámbito privado
// solucionados problemas de errores en GOROOT y GOPATH con
// go mod init
// go mod tidy

func main() {
	var miCoche pk.CarPublic
	miCoche.Marca = "Ferrari"
	miCoche.Año = 2020
	fmt.Println(miCoche)

	// ImprimeMensaje
	pk.ImprimeMensaje("hola")

}
