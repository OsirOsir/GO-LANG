// the use  of Blank Identifier

package main

import "fmt"

func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2

	return
}

func main() {
	length := 36.7
	width := 43.1

	area, _ := rectProps(length, width) //  the blank identifier _ discard the second return value

	fmt.Printf("The area is %.2f\n", area)
}
