package main

import "fmt"

func main() {
	emp3 := struct {
		firstName string
		lastName  string
		age       int
		salary    int
	}{
		firstName: "Philip ",
		lastName:  "Ogutu",
		age:       34,
		salary:    20000,
	}

	fmt.Println("employee3", emp3)

}
