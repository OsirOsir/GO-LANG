package main

import (
	"fmt"
)

func changelocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}
func main() {
	num := [...]int{6, 4, 7, 8, 8}
	fmt.Println("before passing to function", num)
	changelocal(num)
	fmt.Println("after passing to function ", num)
}
