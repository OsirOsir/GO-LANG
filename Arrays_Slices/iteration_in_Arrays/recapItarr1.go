package main

import "fmt"

func main() {
	marks := [5]float64{13.5, 15.3, 100.6, 20.9, 56.6}

	sum := float64(0)

	for i := 0; i < len(marks); i++ {
		fmt.Printf("%d element of marks is %.1f\n", i, marks[i])
		sum += marks[i]
	}
	fmt.Printf("The sume of all this elements is %0.1f\n", sum)
}
