package main

import "fmt"

func main() {
	num := 11

	if num%2 == 0 {
		fmt.Println("The number", num, " provided is even")
		return
	}
	fmt.Println("The number", num, "is odd")
}
