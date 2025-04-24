package main

import "fmt"

func main() {
	a := make([]uint64, 0)
	a = append(a, 1, 2, 3)

	b := append([]uint64{}, a...)
	b = append(b, 4)

	c := append([]uint64{}, a...)
	c = append(c, 5)

	fmt.Println("a:", a) // [1 2 3]
	fmt.Println("b:", b) // [1 2 3 4]
	fmt.Println("c:", c) // [1 2 3 5]
}
