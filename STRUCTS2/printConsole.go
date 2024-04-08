package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	newScanner := bufio.NewReader(os.Stdin)

	fmt.Print("Whats You country? ")

	input, err := newScanner.ReadString('\n')
	if err != nil {
		fmt.Println("You didnt type anything", err)
		return
	}

	input = strings.TrimSpace(input)

	fmt.Println("My country name is ", input)
}
