// 4. Write a program that dynamically allocates memory for an integer
// using the `new` keyword and prints its address and value.

package main

import "fmt"

func main() {
	love := new(int)
	*love = 10
	fmt.Printf("The address is %p and the value is %d\n", love, *love)

	fmt.Println(&love)
}
