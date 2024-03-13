package main

import "fmt"

func main() {
	var names []string
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "Philip", "Joshua", "Eucabeth")
		fmt.Println(names)

	}
}
