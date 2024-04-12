package main

import "fmt"

func factrorialNum(n int) int {
	// result := 1
	// for i := 1; i <= n; i++ {
	// 	result = result * i
	// }

	// return result

	if n == 0 {
		return 1
	}
	return n * factrorialNum(n-1)
}

func main() {
	num := 5

	fmt.Println(factrorialNum(num))
}
