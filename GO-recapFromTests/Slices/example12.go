package main

import "fmt"

func main() {
	fruits := []string{"Mangoe", "Apple"}

	vegetables := []string{"Kales", "Cabbage"}

	food := append(fruits, vegetables...)

	fmt.Println(food)
}
