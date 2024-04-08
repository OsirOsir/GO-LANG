package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	userInsertion := bufio.NewScanner(os.Stdin)

	fmt.Println("Whats your Name: ")

	userInsertion.Scan()

	output := userInsertion.Text()
	if output == "" {
		fmt.Println("You did not type your name")
		return
	}
	fmt.Println("Hello ", output)
}
