package main

import "strings"

func initialLetter(s string) (string, string, string) {
	n := strings.ToUpper(s)
	arr := strings.Split(n, " ")

	var initial []string
	for _, v := range arr {
		initial = append(initial, v[:1])
	}
	if len(initial) >= 3 {
		return initial[0], initial[1], initial[2]
	} else if len(initial) == 2 {
		return initial[0], initial[1], "_"
	} else {
		return initial[0], "_", "_"
	}
}
