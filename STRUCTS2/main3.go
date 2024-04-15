package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "github.com/go-delve/delve/pkg/dwarf/reader"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createInvoice() Invoice {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Create a new invoice name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("Create new Invoice name: ", reader)
	b := newInvoice(name)
	fmt.Println("New invoice name - ", b.name)
	return b
}

func promptOptions(i Invoice) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Chose from a - addItem, t - addTip, s - saveInvoice: ", reader)
	// fmt.Println("You chose ", opt)

	switch opt {
	case "a":
		item, _ := getInput("Add Item: ", reader)
		price, _ := getInput("Add Price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price should be a number")
			promptOptions(i)
		}
		i.updateItems(item, p)
		fmt.Println("Invoice : ", item, price)
		promptOptions(i)

	case "t":
		tip, _ := getInput("Add Tip: ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The Tip should be a number")
			promptOptions(i)
		}
		i.updateTip(t)
		fmt.Println("$", tip)
		promptOptions(i)

	case "s":
		i.save()
		fmt.Println("You have saved the Invoice", i.name)
	default:
		fmt.Println("Invalid choice, chose agin: ")
		promptOptions(i)
	}
}
func main() {
	myInvoice := createInvoice()
	promptOptions(myInvoice)

	// fmt.Println(myInvoice)

}
