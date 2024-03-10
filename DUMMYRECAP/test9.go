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
	sl1 := filterEven([]int{6, 4, 8, 7, 1, 3, 5})
	sl2 := sl1([]int{6, 4, 8, 7, 1, 3, 5})
	fmt.Println(sl2)
}