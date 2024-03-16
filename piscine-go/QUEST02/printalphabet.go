package main

import "fmt"

func main() {
	for i := 'z'; i >= 'a'; i-- {
		fmt.Print(string(i))
	}
	fmt.Println()
}
