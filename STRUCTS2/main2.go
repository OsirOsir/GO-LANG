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
		fmt.Println("Your bill is ", name, price)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		fmt.Println("$", tip)
	case "s":
		fmt.Println("You choose s")
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(i)

	}
}
func main() {
	myInvoice := createInvoice()

	promptOptions(myInvoice)

}
