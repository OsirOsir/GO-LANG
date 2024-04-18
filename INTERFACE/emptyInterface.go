package main

import "fmt"

type i interface{}

func found(val i) {

	fmt.Printf("Type = %T, value = %v\n", val, val)
}

func main() {
	s := "Simplilearn"
	found(s)
}
