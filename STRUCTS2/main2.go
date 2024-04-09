package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func createInvoice() invoice {
	reader := bufio.NewReader(os.Stdin) //  The way we get user input from the consol
	fmt.Print("Create a New Invoice name : ")

	name, _ := reader.ReadString('\n') //
	name = strings.TrimSpace(name)

	i := newInvoice(name)

	fmt.Println("New Invoice created : ", i.name)

	return i
}

func main() {
	myInvoice := createInvoice()

	fmt.Println(myInvoice)
}
