package main

import "fmt"

func main() {
	name := "Alice"
	age := 7

	name = fmt.Sprintf("Hello my name is %s  and I am %d years old", name, age)
	fmt.Println(name)
}
