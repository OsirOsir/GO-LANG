//4. Write a program that dynamically allocates memory for an integer
// using the `new` keyword and prints its address and value.

package main

import "fmt"

func main() {
	a := new(int)
	*a = 1
	fmt.Println(*a)
}
