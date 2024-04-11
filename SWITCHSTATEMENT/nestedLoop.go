package main

import "fmt"

func main() {
	n := 5

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
			// fmt.Println()
		}
		fmt.Println()
	}

	// for i := 0; i < 3; i++ {
	// 	for j := 1; j < 4; j++ {
	// 		fmt.Printf("i = %d , j = %d\n", i, j)
	// 	}
	// }
}
