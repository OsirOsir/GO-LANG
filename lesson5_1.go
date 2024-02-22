package main

import "fmt"

func main() {
	dates := [4]int{23, 45, 67, 56}

	for i, val := range dates {
		fmt.Println(i, val, " ")
	}
	fmt.Println(len(dates))
}
