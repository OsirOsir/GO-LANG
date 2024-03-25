package main

import (
	"fmt"
	"structs/computer"
)

func main() {
	infor := computer.Spec{
		Maker: "apple",
		Price: 500000,
	}
	fmt.Println("Maker:", infor.Maker)
	fmt.Println("Price:", infor.Price)
}
