package main

import "fmt"

func main() {
	myInvoice := newInvoice("Philips Invoice")

	myInvoice.updateItem("Chapo", 34.1)
	myInvoice.updateItem("Beans", 24.1)
	myInvoice.updateItem("Rice", 60.5)
	myInvoice.updateItem("Tea", 10.8)
	myInvoice.updateTip(45.9)

	fmt.Println(myInvoice.format())
}
