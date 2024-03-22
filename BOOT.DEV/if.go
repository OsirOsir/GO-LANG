package main

import "fmt"

func main() {
	email := "hagsksabnmsjakshduwd wnmajl"

	length := len(email)

	fmt.Println(len(email))

	if length < 1 {
		fmt.Println("Email is invalid")
	} else {
		fmt.Println("Email valid")
	}
}
