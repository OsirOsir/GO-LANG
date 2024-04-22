package main

import "fmt"

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printCharacters(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
}

func main() {
	name := "Philip Osir"

	fmt.Printf("String: %s\n", name)
	printBytes(name)
	fmt.Printf("\n")
	printCharacters(name)
	fmt.Printf("\n\n")

	name = "Señor"
	fmt.Printf("String: %s\n", name)
	printCharacters(name)
	fmt.Printf("\n")
	printBytes(name)
	fmt.Println()
}
