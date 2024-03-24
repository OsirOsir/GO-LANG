package main

import "fmt"

func main() {
	love := new(int)

	fmt.Printf("love value is %d, type is %T, address is %v, memory address is %p\n", *love, love, love, &love)

	*love = 56

	fmt.Println("New love value is ", *love)
}
