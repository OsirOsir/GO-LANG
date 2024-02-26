// Question:
// Write a function in Go called sumEven that takes an array of integers as input and returns the sum of all even numbers in the array.

package main

import "fmt"

func sumEven(numbers [17]int) int {

	sum := 0
	for _, v := range numbers {
		if v%2 == 0 {
			sum = v + sum
		}
	}
	return sum
}

func main() {
	numbers := [...]int{6, 1, 2, 5, 6, 5, 9, 8, 2, 1, 7, 5, 7, 9, 1, 8, 5}

	sum := sumEven(numbers)

	fmt.Println(sum)
}
