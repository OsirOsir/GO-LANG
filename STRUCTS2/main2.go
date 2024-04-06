package main

import "fmt"

func main() {
	myInvoice := newInvoice("Hot Meal... Philips Invoice")

	fmt.Println(myInvoice.format())
}
