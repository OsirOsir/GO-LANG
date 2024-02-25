// 6. **Array Search:**
//    - Declare an array named `names` containing 5 strings representing names.
//    - Write a function to search for a specific name in the `names` array and return its index if found, otherwise return -1.

package main

import "fmt"

func searchName(names [5]string, wantedName string) int {
	for i, v := range names {
		if wantedName == v {
			return i
		}
	}
	return -1
}

func main() {
	names := [5]string{"Philip", "Osir", "Iddah", "Awuor", "Alice"}

	names[4] = "Joshua"

	wantedName := "Joshua"

	index := searchName(names, wantedName)
	if index != -1 {
		fmt.Printf("The name %s is available at index %v\n", wantedName, index)
	} else {
		fmt.Printf("The name %s is not available\n", wantedName)
	}
}
