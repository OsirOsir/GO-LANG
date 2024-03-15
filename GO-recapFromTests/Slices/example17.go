package main

import "fmt"

func main() {
	num := []int{9, 8, 6, 8, 6, 9, 9, 8, 6}

	newNum := make([]int, len(num)-3)
	numCopied := copy(newNum, num)
	fmt.Println(numCopied)
	fmt.Println(newNum)

}
