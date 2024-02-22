package main

import "fmt"

func main() {
	for i := 'z'; i >= 'a'; i-- {
		fmt.Printf("%c", i)
	}
	fmt.Println()
}
