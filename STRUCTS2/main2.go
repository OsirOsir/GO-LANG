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

	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createInvoice() invoice { // bellow allows put an input i the console
	reader := bufio.NewReader(os.Stdin) //  The way we get user input from the consol
	// fmt.Print("Create a New Invoice name : ")

	// name, _ := reader.ReadString('\n') // Reads what they are putiing in  until they press on next line
	// name = strings.TrimSpace(name)     //  Trims any  white spaces in the input string put

	name, _ := getInput("Create a new Invoice: ", reader)
	i := newInvoice(name)

	fmt.Println("New Invoice created : ", i.name, " is created")

	return i
}

func promptOptions(i invoice) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option(a - add item, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)
}

func main() {
	myInvoice := createInvoice()
	promptOptions(myInvoice)

}
