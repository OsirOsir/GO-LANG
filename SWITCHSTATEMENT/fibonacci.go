package main

import "fmt"

func fibonacci(n int) {
	a, b := 0, 1

	for i := 0; i <= n; i++ {
		fmt.Print(a, " ")
		a, b = b, (a + b)
	}
	fmt.Println()

}

func main() {
	fibonacci(10)
}
