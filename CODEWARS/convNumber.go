package main

import "fmt"

func Digitize(n int) []int {
	revDig := []int{}

	for n > 0 {
		digit := n % 10

		revDig = append(revDig, digit)
		n /= 10
	}
	return revDig
}

func main() {
	number := 35231
	fmt.Println(Digitize(number))
}
