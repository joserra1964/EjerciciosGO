package main

import "fmt"

func main() {
	var a1 = [3]int{1, 2, 3}
	var a2 = [...]int{4, 5, 6, 7, 8, 9}
	var s1 = [...]string{"Juan", "Pedro", "Luis", "Raquel", "Berto"}
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(s1)
	fmt.Println(s1[0], " y ", s1[3])
	s1[3] = "eva"
	fmt.Println(s1[0], " y ", s1[3])
	var s2 = [5]string{"uno", "dos"}
	fmt.Println(s2, "len : ", len(s2), "cap : ", cap(s2))

	var s3 = [10]string{4: "cuatro", 5: "cinco", 1: "uno"}
	fmt.Println(s3)
	fmt.Println("len : ", len(s3), "cap : ", cap(s3))
}
