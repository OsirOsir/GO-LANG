package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Create a new bill: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create new Invoice ", reader)

	b := newBill(name)
	fmt.Println("New bill created", b)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Select from this Options, a - add Item, s - save, t - addTip :", reader)
	// fmt.Println("Youve selected: ", opt)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item Price: ", reader)
		fmt.Println("Your Bill is", name, price)
	case "s":
		fmt.Println("You chose s")

	case "t":
		tip, _ := getInput("Tip is: ", reader)
		fmt.Println("Your tip is", tip)
	default:
		fmt.Println("That was an invalid choise....")
		promptOptions(b)

	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
	// fmt.Println(myBill)
}
