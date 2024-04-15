package main

import "fmt"

func main() {
	myInvoice := newInvoice("Philips Invoice")

	fmt.Println(myInvoice.format())
}
