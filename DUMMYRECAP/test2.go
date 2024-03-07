package main

import "fmt"

func contains(num int, nums []int) bool {
	for i, v := range nums {
		if v == num {
			fmt.Println("true ", "at index", i)
			return true

		}
	}
	fmt.Println("false")
	return false
}
