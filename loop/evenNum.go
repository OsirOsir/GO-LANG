//Write a program in Go that prints all the even numbers between 1 and 20 using a for loop.

package main

import "fmt"

func main() {
	// for i := 1; i <= 20; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	i := 2

	for i <= 20 {
		fmt.Println(i)
		i += 2
	}
}
