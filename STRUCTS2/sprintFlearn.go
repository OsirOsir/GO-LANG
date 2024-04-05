package main

import "fmt"

func main() {
	name := "Alice\n"
	age := 7

	name = fmt.Sprintf("Hello my name is %s  and I am %d years old", name, age)
	fmt.Println(name)
}
