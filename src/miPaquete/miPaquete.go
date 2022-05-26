package mipaquete

import "fmt"

// CarPublic Car con acceso público
type CarPublic struct {
	Marca string
	Año   int
}

type carPrivado struct {
	marca string
	año   int
}

func ImprimeMensaje(texto string) {
	fmt.Println(texto)
}
