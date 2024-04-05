package main

import (
	"fmt"
	"structs/computer"
)

func main() {
	spec := computer.Spec{
		Maker: "apple",
		Price: 5000,
	}
	fmt.Println("Maker:", spec.Maker)
	fmt.Println("Price:", spec.Price)
}
