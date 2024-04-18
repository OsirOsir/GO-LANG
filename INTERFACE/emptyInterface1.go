package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("Type = %T value = %v\n", i, i)
}

func main() {
	name := "Philip Osir"
	describe(name)
	number := 10
	describe(number)

	p := struct {
		color string
	}{
		color: "red",
	}
	describe(p)

}
