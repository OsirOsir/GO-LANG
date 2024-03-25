package main

import "fmt"

func main() {
	emp3 := struct {
		firstName string
		lastName  string
		age       int
		salary    int
	}{
		firstName: "Kyle",
		lastName:  "Hope",
		age:       18,
		salary:    3000,
	}
	fmt.Println(emp3)
}
