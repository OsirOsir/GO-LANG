package main

import "fmt"

func main() {
	fruits := []string{"Apple", "Orange"}
	vegetables := []string{"Kales", "Tomatoe"}
	food := append(fruits, vegetables...)

	fmt.Println("all the food available are:", food)
}
