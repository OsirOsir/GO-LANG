package main

import "fmt"

func main() {
	myBill := newInvoice("Pan African Hotel Bill")

	myBill.updateItem("Chapati", 30.5)
	myBill.updateItem("Chai", 10.2)
	myBill.updateItem("Beans", 35.8)
	myBill.updateItem("Rice", 20.3)
	myBill.updateItem("Beaf", 28.7)
	myBill.updateTip(36.2)

	fmt.Println(myBill.format())
}
