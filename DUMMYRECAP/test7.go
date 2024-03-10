package main

import "fmt"

func f(x int) int {
	return x * x
}

func applyFunc(f func(int) int, x int) int {
	return f(x)
}

func main() {
	fmt.Println(applyFunc(f, 2))
}
