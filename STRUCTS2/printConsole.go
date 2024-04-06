package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Whats your name: ")
	scanner.Scan()

	input := scanner.Text()

	fmt.Printf("Ypur name is: %q\n", input)
}
