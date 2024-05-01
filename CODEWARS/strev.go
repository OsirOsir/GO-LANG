package main

import "fmt"

func strev(word string) string {
	rev := ""

	for i := len(word) - 1; i >= 0; i-- {
		rev = rev + string(word[i])
	}
	return rev
}

func main() {
	name := "world"

	fmt.Println(strev(name))
}
