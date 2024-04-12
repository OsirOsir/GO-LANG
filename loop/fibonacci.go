package main

import "fmt"

func fibonacci(n int) {
	a, b := 0, 1

	// fmt.Printf("Fibonacci sequence up to th %dth term:\n", n)
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}

func main() {
	fibonacci(10)
}
