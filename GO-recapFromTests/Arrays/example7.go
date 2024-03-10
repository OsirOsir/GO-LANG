package main

import "fmt"

func changeNumber(num [5]int) {
	num[2] = 77
	fmt.Println("modification of the copy", num)
}

func main() {
	num := [5]int{6, 8, 6, 4, 1}
	changeNumber(num)
	fmt.Println("prints the original copy still", num)
}
