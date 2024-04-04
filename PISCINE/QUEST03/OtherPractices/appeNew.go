// 5. Develop a function that takes a pointer to a slice of integers
//and appends a new element to the slice.

package main

import "fmt"

func appNew(a *[]int) {
	// Derefrence the pointer
	*a = append(*a, 89) // Append a new element to the slice
}

func main() {
	// Declare and initialize a slice of integers
	a := []int{85, 87, 87}
	// Pass the address of the slice to the appNew function
	appNew(&a)
	// Print the updated slice
	fmt.Println(a)
}
