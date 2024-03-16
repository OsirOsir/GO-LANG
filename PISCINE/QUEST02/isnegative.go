package main

import "fmt"

func IsNegative(nb int) bool {
	if nb < 0 {
		fmt.Println("T")
		return true
	} else {
		fmt.Println("F")
		return false
	}
}

func main() {
	IsNegative(1)
}
