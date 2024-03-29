package main

import "fmt"

func Swap(a *int, b *int) {
	m := *a
	*a = *b
	*b = m
}

func main() {
	a := 0
	b := 1
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
