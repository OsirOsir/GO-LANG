//3. Implement a function
//that swaps the values of two strings using pointers.

package main

import "fmt"

func swapString(a, b *string) {
	str := *a // Store the value of *a in a str variable
	*a = *b   // Assign the value of *b to *a
	*b = str  // Assign the original value of *a (stored in temp) to *b
}

func main() {
	name1 := "Philip"
	name2 := "Osir"

	swapString(&name1, &name2)
	fmt.Println(name1)
	fmt.Println(name2)
}
