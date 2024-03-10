package main

import "fmt"

func main() {
	a := [5]float64{56.3, 76.2, 66, 90.8, 45}

	for i, v := range a {
		fmt.Printf("the element at index %d is %.1f\n", i, v)
	}

	// for i := 0; i < len(a); i++ {
	// 	fmt.Println("the value at index", i, "is", a[i])
	// 	fmt.Printf("the element at %d is %.1f\n", i, a[i])
	// }
}
