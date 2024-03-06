package main

import "fmt"

func main() {
	age := 54

	if age := 54; age <= 50 {
		fmt.Printf("%v is young\n", age)
		return
	}
	fmt.Printf("%v is older\n", age)
}
