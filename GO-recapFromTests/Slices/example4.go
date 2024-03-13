package main

import "fmt"

func main() {
	numa := [3]int{78, 79, 80}

	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("Orininal array before any change", numa)

	nums1[0] = 100
	fmt.Println("array after first modification of nums1 at index 0", numa)

	nums2[1] = 101
	fmt.Println("array after first modification of nums2 at index 1", numa)
}
