// 2. Create a function that accepts a slice of integers and returns
// a pointer to the maximum value in the slice.

package main

import "fmt"

func sliceInt(x []int) *int {

	maximum := 0
	for _, i := range x {
		if i > maximum {
			maximum = i
		}
	}
	return &maximum
}

func main() {
	a := []int{4, 5, 6, 5, 7, 9}

	maxNew := sliceInt(a)
	fmt.Println(*maxNew)
}
