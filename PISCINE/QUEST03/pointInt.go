//1. Write a function that takes a pointer to an integer and increments its value by 1.

package main

import "fmt"

func pointInt(a *int) {
	*a = *a + 1

}

func main() {
	x := 1
	pointInt(&x)
	fmt.Println(x)
}
