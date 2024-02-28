package main

import "fmt"

func main() {
	fruits := [][]string{
		{"mangoes", "Apples"},
		{"orange", "grapes"},
		{"berries", "lemon"},
	}
	food := fruits
	food[2] = []string{"Tomato", "Onions"}
	fmt.Println(food, len(fruits))
}
