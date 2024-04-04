// Implement a function that swaps the values of two strings using pointers.
package main

import "fmt"

func swap(a, b *string) {
	temp := *a
	*a = *b
	*b = temp

}

func main() {
	a := "Philip"
	b := "Otieno"

	swap(&a, &b)

	fmt.Println(a)
	fmt.Println(b)

}
