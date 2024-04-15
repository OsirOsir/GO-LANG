package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	name, err := r.ReadString('\n')

	return strings.TrimSpace(name), err
}

func createInvoice() invoice {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Create a new Invoice name: ")

	// name, _ := reader.ReadString('\n')

	// name = strings.TrimSpace(name)
	name, _ := getInput("Create a new Invoice Name: ", reader)

	i := newInvoice(name)
	fmt.Println("New Invoice Created", name)

	return i
}

func promptOptions(i invoice) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Chose from a - addItem, t - addTip, s - SaveInvoice ", reader)

	// fmt.Println("You Chose :", opt)

	switch opt {
	case "a":
		item, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item Price: ", reader)

		fmt.Println("Bill is :", item, price)
		promptOptions(i)

	case "t":
		tip, _ := getInput("Add tip: ($)", reader)

		fmt.Println("$", tip)
		promptOptions(i)

	case "s":
		fmt.Println("Bill saved")

	default:
		fmt.Println("Your choice is invalid")
		promptOptions(i)
	}
}

func main() {
	myInvoice := createInvoice()
	promptOptions(myInvoice)
}
