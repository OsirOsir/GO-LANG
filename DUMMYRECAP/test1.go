package main

import "fmt"

func find(num int, nums ...int) {
	fmt.Printf("nums is of type  %T\n", nums)

	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "is available at index", i, "in ", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
}
