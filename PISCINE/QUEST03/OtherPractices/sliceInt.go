// Create a function that accepts a slice of integers and returns a pointer
// to the maximum value in the slice.

package main

import "fmt"

func slicInt(a []int) *int {
	max := a[0]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return &max
}

func main() {
	a := []int{8, 9, 6, 8, 9, 6, 8, 6}

	maximum := slicInt(a)
	fmt.Println(*maximum)

}
