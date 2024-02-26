package main

import "fmt"

func main() {
	ages := []int{56, 15, 26, 45, 97, 66}
	ages[5] = 43
	ages = append(ages, 65)

	fmt.Println(ages, len(ages))

	//range in slices

	rangOne := ages[1:]
	rangTwo := ages[1:3]
	rangThree := ages[:4]

	fmt.Println(rangOne, rangTwo, rangThree)
}
