package main

import "fmt"

/*	slice parte
	nombre := []datatype{values}
	miSlace:=[]int
*/

func main() {
	sl1 := []int{}
	fmt.Println(sl1, "len : ", len(sl1), "cap : ", cap(sl1))

	sl2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(sl2, "len : ", len(sl2), "cap : ", cap(sl2))

	// creando un slice troceando un array
	arr1 := [6]int{11, 12, 14, 16, 1, 6}
	sl3 := arr1[2:4]
	fmt.Println(sl3, "len : ", len(sl3), "cap : ", cap(sl3))

	// add a un slice
	sl3 = append(sl3, arr1[0], arr1[5], arr1[0])
	fmt.Println(sl3, "len : ", len(sl3), "cap : ", cap(sl3))

	sl3 = append(sl3, 0, -14, -15)
	fmt.Println(sl3, "len : ", len(sl3), "cap : ", cap(sl3))

	sl4 := []int{111, 122, 133, 44}
	sl3 = append(sl3, sl4...)
	fmt.Println(sl3, "len : ", len(sl3), "cap : ", cap(sl3))

	sl5 := []int{}
	sl3 = append(sl5, 1, 2, 3)
	fmt.Println(sl3, "len : ", len(sl3), "cap : ", cap(sl3))

	// copiando a un slice
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Printf("numeros = %v, len = %d, cap = %d \n", numeros, len(numeros), cap(numeros))
	necesarios := numeros[:len(numeros)-10]
	copia := make([]int, len(necesarios))
	copy(copia, necesarios)
	fmt.Printf("copia = %v, len = %d, cap = %d \n", copia, len(copia), cap(copia))

}
