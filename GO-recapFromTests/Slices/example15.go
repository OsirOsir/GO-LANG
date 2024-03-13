package main

import "fmt"

func main() {
	num := []int{8, 9, 8, 7, 6, 5, 5}

	for i, v := range num {
		fmt.Println("The index of value", v, "is", i)
	}

	for i := range num {
		fmt.Println("The indices are ", i)
	}

	for v := 0; v < len(num); v++ {
		fmt.Printf("the value of index %d is %d\n", v, num[v])
	}
}
