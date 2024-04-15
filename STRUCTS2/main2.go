package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// helper function
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createInvoice() invoice {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Creat a new invoice name: ")
	// name, _ := reader.ReadString('\n')

	// name = strings.TrimSpace(name)
	name, _ := getInput("Creat a new invoice name: ", reader)

	i := newInvoice(name)

	fmt.Println("New Invoice Created ", i.name)

	return i

}

func promptOptions(i invoice) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save invoice, t - add tip): ", reader)
	switch opt {
	case "a":
		name, _ := getInput("item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be  anumber ")
			promptOptions(i)
		}
		i.updateItem(name, p)
		fmt.Println("Item added ", name, price)
		promptOptions(i)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The Tip must be a number")
			promptOptions(i)
		}
		i.updateTip(t)
		fmt.Println("Tip added :  $", tip)
		promptOptions(i)
	case "s":
		i.save()
		fmt.Println("You saved ", i.name)
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(i)

	}
}
func main() {
	myInvoice := createInvoice()

	promptOptions(myInvoice)

}
