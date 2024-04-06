package main

import "fmt"

func main() {

	myInvoice := newInvoice("Hotel Bill")

	myInvoice.updateItem("Chapati", 45.5)
	myInvoice.updateItem("Tea", 10.5)
	myInvoice.updateItem("Beans", 35.9)
	myInvoice.updateItem("GreenGrams", 34.6)
	myInvoice.updateItem("Beef", 40.2)
	myInvoice.updateItem("Matumbo", 65.1)
	myInvoice.updateTip(23.6)

	fmt.Println(myInvoice.format())
}
