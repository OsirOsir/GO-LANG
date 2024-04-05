package main

import "fmt"

type image struct {
	name string
	dime map[int]int
}

func main() {
	image1 := image{
		name: "Image Dimensions",
		dime: map[int]int{
			34: 20,
		},
	}

	fmt.Println(image1.name)

	for k, v := range image1.dime {
		fmt.Printf("Width: %dcm Hight: %vcm \n", k, v)
	}

	// image2 := image{
	// 	name: "Image Dimensions",
	// 	dime: map[int]int{
	// 		30: 40,
	// 	},
	// }

	// if image1 == image2 {
	// 	fmt.Println("the two Images are equal")
	// } else {
	// 	fmt.Println("the two Images are not equal")
	// }

	// .\mapsNotComparable.go:31:5: invalid operation: image1 == image2
	// (struct containing map[int]int cannot be compared)
}
