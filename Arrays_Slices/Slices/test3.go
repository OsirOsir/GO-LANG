package main

import "fmt"

func main() {
	digits := [...]int{57, 58, 90, 82, 100, 78, 69, 67, 59}

	digUpdate := digits[2:5] // Slice digits from index 2 up to, but not including, index 5
	fmt.Println(digits)
	for v := range digUpdate { // Iterate over the indices of digUpdate
		digUpdate[v]++ // Increment each element of digUpdate by one
	}
	fmt.Println(digits) // Print the updated digits slice
}
