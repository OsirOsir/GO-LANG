package main

import "fmt"

func main() {
	a := 45

	var b *int
	if b == nil {
		fmt.Println("b is", b)
		fmt.Println(&a)
	}
	b = &a
	fmt.Printf("b after initialization, the momory address of a is %p\n", b)
}
