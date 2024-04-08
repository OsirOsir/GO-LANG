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

	input, _ := newScanner.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("An eeror occurred:", err)
	// 	return
	// }

	input = strings.TrimSpace(input)

	if input == "" {
		fmt.Println("You didnt type anything.")
		return
	}

	fmt.Println("My country name is ", input)
}
