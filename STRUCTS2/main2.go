package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// helper function
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n') // get the information

	return strings.TrimSpace(input), err
}

func createInvoice() invoice { // bellow allows put an input i the console
	reader := bufio.NewReader(os.Stdin) //  The way we get user input from the consol
	// fmt.Print("Create a New Invoice name : ")

	// name, _ := reader.ReadString('\n') // Reads what they are putiing in  until they press on next line
	// name = strings.TrimSpace(name)     //  Trims any  white spaces in the input string put

	name, _ := getInput("Create a new invoice name: ", reader)
	i := newInvoice(name)

	fmt.Println("New Invoice created : ", i.name, " is created")

	return i
}

func main() {
	myInvoice := createInvoice()

	fmt.Println(myInvoice)
}
