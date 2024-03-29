package main

import "fmt"

func StrLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}
