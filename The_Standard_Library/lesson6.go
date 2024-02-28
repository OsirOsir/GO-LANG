package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := []int{7, 8, 2, 1, 8, 6, 1, 8, 5}
	agesR := ages
	sort.Sort(sort.Reverse(sort.IntSlice(agesR)))
	// fmt.Println(agesR)
	for i, v := range agesR {
		fmt.Println(v, i)
	}

	// sort.Ints(ages)

	// fmt.Println(ages)
}
