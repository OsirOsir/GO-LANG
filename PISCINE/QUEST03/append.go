//5. Develop a function that takes a pointer to a slice of integers and appends a
//new element to the slice.

package main

import "fmt"

func appenSlice(a *[]int) []int {
	*a = append(*a, 90)

	return *a

}

func main() {
	a := []int{98, 78, 89}
	appenSlice(&a)
	fmt.Println(a)
}
