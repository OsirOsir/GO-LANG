package main

import (
	"fmt"
	"strings"
)

func initLetter(s string) (string, string, string) {
	n := strings.ToUpper(s)
	arr := strings.Split(n, " ")

	var initials []string
	for _, v := range arr { //So, the loop iterates over each word in the arr slice, and for each word, it extracts the first character (initial) and appends it to the initials slice.
		initials = append(initials, v[:1])
	}
	if len(initials) >= 3 {
		return initials[0], initials[1], initials[2]
	} else if len(initials) <= 2 {
		return initials[0], initials[1], "_"
	} else {
		return initials[0], "_", "_"
	}
}

func main() {
	s := "Iddah Awuor Otieno"
	in1, in2, in3 := initLetter(s)

	fmt.Println(in1, in2, in3)
}
