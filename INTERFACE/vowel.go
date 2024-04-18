package main

import "fmt"

type vowel interface {
	findVowel() []rune
}

type MyString string

func (ms MyString) findVowel() []rune {
	v := []rune{}

	for _, r := range ms {
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
			v = append(v, r)
		}
	}
	return v
}

func main() {
	name := MyString("Iddah Awuor Otieno")

	var vl vowel

	vl = name

	fmt.Printf("vowels are %c\n", vl.findVowel())
}
