package main

import "fmt"

func main() {
	dates := [4]int{23, 45, 67, 56}

	for _, val := range dates {
		fmt.Println(val, " ")
	}
	fmt.Println(len(dates))
}
