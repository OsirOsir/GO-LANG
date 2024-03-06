package main

import (
	"fmt"
	"strings"
)

func main() {
	// Define a string to split
	str := "philip"

	// Split the string into words using space as the delimiter
	s := strings.ToUpper(str)
	arr := strings.Split(s, " ")
	var initials []string
	for _, v := range arr {
		initials = append(initials, v[:1])
	}
	if len(initials) > 1 {
		fmt.Println(initials[0], initials[1])
	} else {
		fmt.Println(initials[0], "_")
	}
}
