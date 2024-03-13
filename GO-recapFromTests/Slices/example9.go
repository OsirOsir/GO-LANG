package main

import "fmt"

func main() {
	b := make([]int, 7, 7)

	b[0] = 1
	b[1] = 2
	b[2] = 3
	b[3] = 4
	b[4] = 5
	b[5] = 6
	b[6] = 7

	fmt.Println("Original slice of b is", b, "and capacity is", cap(b))

	e := b[2:4]
	fmt.Println("Slicing of b", e, "and Capacity is ", cap(e))

	// reslicing
	e = e[:cap(e)]
	fmt.Println("Reslicing of e is ", e, cap(e), "length", len(e))
}
