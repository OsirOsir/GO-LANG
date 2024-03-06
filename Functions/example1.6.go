// Multiple Return Values  and also giving returns names

package main

import "fmt"

func rectProps(length, width float64) (float64, float64) {
	area := length * width
	perim := (length + width) * 2

	return area, perim
}

func main() {
	length, width := 20.8, 28.5

	area, perim := rectProps(length, width)

	fmt.Printf("The area  is %.1f while the perimeter is %.2f\n", area, perim)

}

// If a function returns multiple return values then they must be specified between ( and ).
// func rectProps(length, width float64)(float64, float64) has two float64 parameters
// length and width and also returns two float64 values.
