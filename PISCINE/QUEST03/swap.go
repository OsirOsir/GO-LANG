package main

import "fmt"

func Swap(a *int, b *int) {
	m := *a

	*a = *b
	*b = m
}

func main() {
	a := 1
	b := 0

	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
