package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	newScanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Whats You country? ")

	newScanner.Scan()

	country := newScanner.Text()

	fmt.Println("My country name is ?", country)
}
