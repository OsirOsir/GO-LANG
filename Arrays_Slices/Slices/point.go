package main

import "fmt"

func PopInt(ints []int) []int {
	if len(ints) == 0 {
		return []int{}
	}
	return ints[:len(ints)-1]
}
func main() {
	fmt.Println(PopInt([]int{6, 7, 8, 9}))
}
