package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "Hello there Guys!"
	// intro := "How are you ninjas"
	ages := []int{61, 25, 81, 53, 16, 55, 68, 15}
	sort.Ints(ages)
	fmt.Println(ages)
	//  index := sort.SearchInts(ages, 81)
	fmt.Println(sort.SearchInts(ages, 81))
	// fmt.Println(index)
	// fmt.Println(strings.Contains(greeting, "Hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "Bigup"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Split(intro, " "))
	// fmt.Println(strings.Index(greeting, "y"))

	//  The original value  is unchanged
	// fmt.Println("original string value = ", greeting)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names)

	fmt.Println(names, len(names))
	// index := sort.SearchStrings(names, "mario")

	fmt.Println(sort.SearchStrings(names, "bowser"))
}
