package main

import "fmt"

func main() {
	for i := 'A'; i <= 'Z'; i++ {
		if i%2 != 0 {
			fmt.Print(string(i))
		} else {
			fmt.Print(string(i + 32))
		}
	}
	fmt.Println()
}
