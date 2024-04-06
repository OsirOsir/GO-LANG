package main

import "fmt"

func main() {
	myInvoice := newInvoice("Hot Meal... Philips Invoice")

	myInvoice.updateItems("Chapatis", 40.5)
	myInvoice.updateItems("Beans", 84.5)
	myInvoice.updateItems("Tea", 24.5)
	myInvoice.updateItems("Ugali", 10.5)
	myInvoice.updateTip(10.1)

	fmt.Println(myInvoice.format())
}
