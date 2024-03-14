package main

import "fmt"

func strLn(s string) int {
	n := 0
	for range s {
		n++
	}
	return n
}

func main() {

	s := "Hello world"

	fmt.Println(strLn(s))
}
