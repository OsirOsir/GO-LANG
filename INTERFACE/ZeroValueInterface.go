package main

import "fmt"

type Describer interface {
	Describe()
}

func main() {
	var d1 Describer
	if d1 == nil {
		fmt.Printf("d1 is nil and has a type %T value %v\n", d1, d1)
	}
}
