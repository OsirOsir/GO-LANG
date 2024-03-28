package main

import "fmt"

func PrintNbr(n int) {
	fmt.Print(n)
}
func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	fmt.Printf("\n")
}
