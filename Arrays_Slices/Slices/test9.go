// Appending to a slice

package main

import "fmt"

func main() {
	var names []string // zero value of slice is nill
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sabastian", "Vinay")
		fmt.Println("New names content:", names)
	}
}
