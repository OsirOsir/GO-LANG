// 7. **Array Sorting:**
//    - Declare an array named `numbers` containing 5 integers in random order.
//    - Write a program to sort the `numbers` array in ascending order and print the sorted array.

package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := [...]int{7, 69, 80, 13, 9, 81, 64, 8, 72, 1, 5, 2, 7, 8, 2, 1}

	sort.Ints(numbers[:]) // This method directly sorts the array numbers  in place, It sorts ellements of numbers in an ascending order modifying the original array
	fmt.Println(numbers)
}

//   sortNum := numbers[:]    this creates a slice sortNum that refers to the entire array of numbers  then
//   sort.Ints(sortNum)     sorts this slices in place, while the slice itself is sorted  the original array numbers remain an changed
// this means that any changes made to sortNum do not affect numbers and Vice versa
