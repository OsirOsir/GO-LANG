package main

import "fmt"

type address struct {
	country  string //  country and Province plus localadd are promoted fields
	province string
	localadd int
}
type personalDetails struct {
	name   string
	age    int
	status string
	address
}

func main() {
	bio := personalDetails{
		name:   "Philip Osir Otieno",
		age:    30,
		status: "single",
		address: address{
			country:  "Kenya",
			province: "Nyanza",
			localadd: 40100,
		},
	}

	fmt.Printf("The country of Philip is %s and Province %s, while the addressNo. is %d\n:", bio.country, bio.address.province, bio.address.localadd)
}
