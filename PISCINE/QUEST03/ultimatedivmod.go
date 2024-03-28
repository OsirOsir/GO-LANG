package main

import "fmt"

func UltimateDivMod(a *int, b *int) {
	div := *a / *b
	mod := *a % *b

	*a = div
	*b = mod
}

func main() {
	a := 13
	b := 2

	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
