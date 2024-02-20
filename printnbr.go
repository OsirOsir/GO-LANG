package main

import "fmt"

func PrintNbr(n int) {
	fmt.Print(n)

	fmt.Println('\n')
}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	fmt.Print('\n')
}
