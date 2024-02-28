package main

import "fmt"

func main() {
	grades := []int{92, 88, 95}

	// grades[2] = 97

	for v := range grades {
		fmt.Println(grades[v])
	}

	// middlGrades := grades[1:3]

	// fmt.Println(middlGrades)
	// fmt.Println(grades[1])
	updGrades := append(grades, 100, 80, 85)
	fmt.Println(updGrades, len(grades), cap(grades))
	lessGrades := grades[1:]

	fmt.Println(lessGrades)

}
