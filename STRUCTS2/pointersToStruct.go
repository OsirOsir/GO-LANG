package main

import "fmt"

type vendorDetails struct {
	name   string
	goods  string
	price  float64
	profit int
}

func main() {
	hust := &vendorDetails{
		name:   "Eucabeth",
		goods:  "vegetables",
		price:  30.5,
		profit: 5,
	}

	fmt.Println("Name of the Vendor is ", (*hust).name) // sytntax is (*hust).name  but we can still use hust.name to still acces the name
	fmt.Println("Type of Goods she sells is ", (*hust).goods)

}
