// Question:
// Write a function called filterEven that takes a slice of integers and returns a function.
// The returned function should filter out even numbers from the input slice and return
// a new slice containing only the odd numbers.

package main

import "fmt"

func filterEven(x []int) func([]int) []int {
	return func(y []int) []int {
		var result []int
		for _, v := range y {
			if v%2 != 0 {
				result = append(result, v)
			}
		}
		return result
	}
}

func main() {

	in1 := filterEven([]int{3, 4, 5, 6, 1, 9})

	in2 := in1([]int{3, 4, 5, 6, 1, 9})

	fmt.Println(in2)
}
