package main

import "fmt"

func main() {
	a := [3]int{2, 4, 2}
	b := [6]int{4, 5, 3, 4, 6}
	a = b
	fmt.Println(b, a)
}

// This is not possible since you cannot use b (variable of type [6]int) as [3]int value in assignment    not unless the sizes are the same

// Pro Osir@Philip-Osir-PC:Arrays$ go run example5.go
// # command-line-arguments
// .\example5.go:8:6: cannot use b (variable of type [6]int) as [3]int value in assignment
