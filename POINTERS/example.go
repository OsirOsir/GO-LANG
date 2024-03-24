package main

import "fmt"

func main() {
	b := 125

	var a *int = &b

	fmt.Printf("the address of b is %p\n", &b)
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("adress of b is", a)

}
