package main

import "fmt"

func main() {
	pls := [2][2]string{
		{"C", "C++"},
		{"JavaScript", "Java"},
		{"Go", "Rust"},
	}

	for _, r := range pls {
		for _, c := range r {
			fmt.Printf("%s ", c)
		}
		fmt.Printf("\n")
	}
}
