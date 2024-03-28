package main

import "fmt"

func PrintStr(s string) {

	for _, i := range s {
		fmt.Print(string(i))
	}
	fmt.Println()
}
func main() {
	PrintStr("Hello World!")
}
