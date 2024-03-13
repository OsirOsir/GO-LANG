package main

import "fmt"

func main() {
	a := make([]int, 4, 4)

	fmt.Println("initial slice", a)

	a[0] = 1
	a[1] = 3
	a[2] = 4
	a[3] = 5

	fmt.Println("modified slice", a)

	b := a
	b[3] = 7
	fmt.Println("Final Modified ", b)

	b = append(b, 6)
	fmt.Println("appended", b)

	//slicing the slice

	c := b[1:3]
	fmt.Println("slices", c)

	fmt.Println("final plus appended", b, "and the  length is: ", len(b), "while the capacity is", cap(b))

	// reslicing slices

	e := b[2:4]
	fmt.Println(e, cap(e))

	e = e[:cap(e)]
	fmt.Println(e)

}
