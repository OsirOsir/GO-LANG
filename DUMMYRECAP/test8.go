// Anonymous Function:

//Write a function called incrementBy that takes an integer x and returns a function
//that increments its input by x.

package main

import "fmt"

func incrementBy(x int) func(int) int {
	return func(y int) int {
		return y + x
	}
}

func main() {
	inc1 := incrementBy(2)

	result := inc1(5)

	fmt.Println(result)
}