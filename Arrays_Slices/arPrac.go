package main

import "fmt"

func changNumes(num [8]int) {
	num[5] = 40
	for _, val := range num {
		fmt.Println(val)
	}
}

func main() {
	num := [...]int{45, 87, 826, 6, 62, 56, 85, 34}
	changNumes(num)
}
