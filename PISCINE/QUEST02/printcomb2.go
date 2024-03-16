package main

import "fmt"

func PrintComb2() {
	for x := 0; x <= 98; x++ {
		for y := x + 1; y <= 99; y++ {
			fmt.Printf("%.2d %.2d", x, y)
			if x != 98 || y != 99 {
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
