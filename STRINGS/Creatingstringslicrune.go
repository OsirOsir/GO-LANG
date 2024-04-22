package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// runes := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}

	// str := string(runes)
	// fmt.Println(str)

	name := "philipojkhAIKJj"
	runeSlice := []rune(name)

	fmt.Printf("%c %d\n", runeSlice, utf8.RuneCountInString(name))
}
