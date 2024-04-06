package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// creating a new scanner to read input from statndard input (console)
	scanner := bufio.NewScanner(os.Stdin)
	// Printing a message to the console, prompting the user to type something
	fmt.Printf("Type something: ")

	//waiting for the user to input something and press Enter
	scanner.Scan()

	// Retreiving the input provide by the user
	input := scanner.Text()

	//Printing out what the user typed , surrounded by double quotes
	fmt.Printf("You typed: %q\n", input)
}
