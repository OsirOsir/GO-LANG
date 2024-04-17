package main

import "fmt"

var i interface{}

func found(i interface{}) {

	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "Simplilearn"
	found(s)
}
