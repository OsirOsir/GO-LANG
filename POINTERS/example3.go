package main

import "fmt"

func main() {
	love := new(int)

	fmt.Printf("size value is %d, type is %T, address is %v, memory address is %p\n", *love, love, love, &love)
}
