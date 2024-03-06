package main

import (
	"fmt"
	"strings"
)

func initLetter(s string) (string, string) {
	arr := strings.Split(s, " ")
	initial := make([]string, 0)
	for _, v := range arr {
		initial = append(initial, v[:1])
	}
	if len(initial) > 1 {
		return initial[0], initial[1]
	}
	return initial[0], "_"
}

func main() {
	s := "Alice Bob"
	in1, in2 := initLetter(s)

	fmt.Println(in1, in2)
}
