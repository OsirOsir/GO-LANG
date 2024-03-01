package main

import "fmt"

func main() {
	for x := 'a'; x <= 'z'; x++ {
		if x%2 == 0 {
			fmt.Print(string(x))
		}
	}
	fmt.Println()
}
