package main

import (
	"fmt"
	"structs1/structFile"
)

func main() {
	infor := structFile.Phone{
		Name:  "Tecno",
		Size:  5.5,
		Model: "S5",
	}

	fmt.Println("Name of Phone is :", infor.Name)
}
