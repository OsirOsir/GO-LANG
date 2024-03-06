package main

import "fmt"

func main() {
	for i, j := 10, 1; i <= 19 && j <= 10; i, j = i+1, j+1 {
		fmt.Printf("%d * %d = %d\n", i, j, i*j)
	}
}
