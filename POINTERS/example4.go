package main

import "fmt"

func main() {
	b := 255

	a := &b

	fmt.Println("address of b is ", a)
	fmt.Println("value of b is ", *a)
	fmt.Printf("The typ of b is %T and the adress of a %p\n ", a, &a)
}
