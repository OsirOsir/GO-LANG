package main

import "fmt"

func main() {
	n, m := 12, 5
	quot, rem := divide(n, m)

	fmt.Printf("The quotient is %d while the remainder is %d\n", quot, rem)
}
