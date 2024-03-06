// continue and break
package main

import "fmt"

func main() {
	names := []string{"Patrick", "Alice", "Iddah", "Yuanita", "Kyle", "Benard"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continue after index", index)
			continue
		}
		if index > 3 {
			fmt.Println("Break at index ", index)
			break
		}
		fmt.Println("The index of ", value, "is ", index)
	}

}
