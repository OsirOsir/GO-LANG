package main

import "fmt"

func main() {
	num := 231

	if num%2 == 0 {
		fmt.Println("The number ", num, "is even")
	} else {
		fmt.Println("The number ", num, "is odd")
	}
}
