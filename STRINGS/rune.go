package main

import "fmt"

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChar(s string) {
	fmt.Printf("Characters: ")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}
}

func main() {
	name := "Hello guys"
	printBytes(name)
	fmt.Println()

	printChar(name)
	fmt.Println()

}
