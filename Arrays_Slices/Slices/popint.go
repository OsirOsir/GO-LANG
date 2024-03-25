package main

import "fmt"

func popint(x []int) []int {
	if len(x) == 0 {
		return []int{}
	}
	return x[:len(x)-1]
}

func main() {
	y := []int{98, 68, 75, 98, 78}

	fmt.Println(popint(y))

}
