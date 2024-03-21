package main

import "fmt"

func PrintComb2() {
	for y := 0; y <= 98; y++ {
		for z := y + 1; z <= 99; z++ {
			fmt.Printf("%.2d ", y)
			fmt.Printf("%.2d", z)
			if y <= 98 || z <= 99 {
				fmt.Print(",")
				fmt.Print(" ")
			}
		}
	}
	fmt.Println()
}
func main() {
	PrintComb2()
}
