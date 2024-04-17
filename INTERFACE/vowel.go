package main

import "fmt"

type VowelFinder interface {
	findVowel() []rune
}

type myVowel string

func (ms myVowel) findVowel() []rune {
	vowel := []rune{}

	for _, r := range ms {
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
			vowel = append(vowel, r)
		}
	}
	return vowel
}

func main() {
	name := myVowel("Philip Osir Otieno")

	var vl VowelFinder

	// vl = name

	fmt.Printf("The vowels are %c \n", vl.findVowel())
}
