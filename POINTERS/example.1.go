package main

import "fmt"

func main() {
	var x string = "45"

	var y *string = &x
	fmt.Println("Memory address of x is ", &x)
	fmt.Printf("the memory address of x is : %p\n", y)
}
