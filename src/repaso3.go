package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var (
	maxint uint64 = 1<<64 - 1
	i2            = 1 << 3
	i3            = 2 << 3
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func main() {

	fmt.Println("Mi número favorito es : ", rand.Intn(100))
	fmt.Println("Ahora es ", time.Now())
	fmt.Println("El resultado %g.", math.Sqrt(7))
	fmt.Println(maxint, " - ", i2, " - ", i3)
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
