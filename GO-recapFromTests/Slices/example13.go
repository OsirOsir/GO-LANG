package main

import "fmt"

func negValues(n []int) {

	for i := range n {
		n[i] -= 2
	}
}

func main() {
	n := []int{6, 5, 9, 6, 5}

	negValues(n)
	fmt.Println(n)
}
