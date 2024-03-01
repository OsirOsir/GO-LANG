// Nested loops

package main

import "fmt"

func main() {
	n := 5 //  we want to print 5 line

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
