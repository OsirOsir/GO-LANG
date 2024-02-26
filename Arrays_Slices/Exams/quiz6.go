// 6. **Array Search:**
//   - Declare an array named `names` containing 5 strings representing names.
//   - Write a function to search for a specific name in the `names` array and return its index if found, otherwise return -1.

package main

import "fmt"

func searchName(names []string, targetName string) int {

	for i, v := range names {
		if targetName == v {
			return i
		}
	}
	return -1
}

func main() {
	names := []string{"Joshua", "Alex", "Vincent", "Caleb", "Mpangala", "Alice"}
	targetName := "Alex"
	index := searchName(names, targetName)
	if index != -1 {
		fmt.Printf("The name %s is available at index %v\n", targetName, index)
	} else {
		fmt.Printf("The name %s is not found\n", targetName)
	}

	// fmt.Println(index)
}
