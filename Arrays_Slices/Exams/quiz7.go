// 7. **Array Sorting:**
//    - Declare an array named `numbers` containing 5 integers in random order.
//    - Write a program to sort the `numbers` array in ascending order and print the sorted array.

package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := [...]int{67, 54, 38, 85, 45}
	sortNumbers := numbers[:] // converts arrays to slices using a slice expression

	sort.Ints(sortNumbers) // sorts the slices in an assending order
	fmt.Println(sortNumbers)

	sort.Sort(sort.Reverse(sort.IntSlice(sortNumbers))) // sorts the slices in an descending order
	fmt.Println(sortNumbers)
}

//sort.Float64s()  .......... his is for floats ascending.................... sort.Sort(sort.Reverse(sort.Float64Slice(sortNumbers))) discending
