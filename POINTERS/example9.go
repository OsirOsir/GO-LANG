package main

import "fmt"

func modify(sls []int) {
	sls[1] = 100
}

func main() {
	num := []int{65, 87, 48, 89}

	modify(num[:])

	fmt.Println(num)
}
