package main

import "fmt"

func reverseElement(letters []rune) {
	for i, j := 0, len(letters)-1; i < j; i, j = i+1, j-1 {
		letters[i], letters[j] = letters[j], letters[i]
	}
}
func main() {
	letters := []rune{'P', 'H', 'I', 'L', 'I', 'P'}
	reverseElement(letters)

	fmt.Println(string(letters))
}
