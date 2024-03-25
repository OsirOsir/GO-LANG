package main

import "fmt"

func change(nName *[3]string) {
	(*nName)[1] = "Davin"
}

func main() {
	name := [3]string{"Eucabet", "Kyle", "Robert"}

	change(&name)
	fmt.Println(name)
}
