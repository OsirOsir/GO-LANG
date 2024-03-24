package main

import "fmt"

func hello() *int {
	i := 5
	return &i
}

func main() {
	b := hello()
	fmt.Println("Value of b", *b)
}
