package main

import "fmt"

func factorial(n int) int {
	// result := 1

	// for i := 1; i <= n; i++ {
	// 	result = result * i
	// }
	// return result

	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(5))
}
