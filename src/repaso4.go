package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true}, {3, false}, {5, true}, {7, true}, {11, false}, {13, true},
	}

	fmt.Printf("len=%d  cap=%d  %v\n", len(s), cap(s), s)

	t := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("len=%d  cap=%d  %v\n", len(t), cap(t), t)
	t = t[1:4]
	fmt.Printf("len=%d  cap=%d  %v\n", len(t), cap(t), t)
	t = t[:2]
	fmt.Printf("len=%d  cap=%d  %v\n", len(t), cap(t), t)
	t = t[1:]
	fmt.Printf("len=%d  cap=%d  %v\n", len(t), cap(t), t)

}
