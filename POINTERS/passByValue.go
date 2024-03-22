package main

import "fmt"

func updateName(x string) {
	x = "wedge"
}

func main() {
	name := "tifa"

	updateName(name)
	fmt.Println(name)
	fmt.Println("Memory address of name is ", &name)

	m := &name

	fmt.Println("memory address", m)
	fmt.Println("The value at memory addrees", *m)

}

// Exaplanation
/*
|--name--|---m-----|
| 0x001  |  0x002  |
|--------|---------|
| "tifa" |  p0x001 |
|--------|---------|
*/
