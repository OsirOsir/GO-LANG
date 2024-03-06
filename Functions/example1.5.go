package main

import "fmt"

func areaCalculation(l int, w int) int {
	totalArea := l * w

	return totalArea
}

func main() {
	l, w := 34, 45

	totalArea := areaCalculation(l, w)
	fmt.Println("The area of  the Rectangle is:", totalArea, "cm")
}
