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
	fmt.Println("Hello ", output)
}
