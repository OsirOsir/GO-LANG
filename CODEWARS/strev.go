package main

import "fmt"

func strev(word string) string {
	rev := ""
	for _, v := range word {
		rev += string(v)
	}
	return rev
}

func main() {
	name := "world"

	fmt.Println(strev(name))
}
