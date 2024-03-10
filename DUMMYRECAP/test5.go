package main

import "fmt"

func divide(n, m int) (int, int) {
	qut := n / m
	rem := n % m
	fmt.Println(qut, rem)
	return qut, rem
}

func main() {
	qut := 5
	rem := 2
	divide(qut, rem)
}
