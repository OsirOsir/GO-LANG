//6. Create a function that accepts a pointer to a
// structure representing a student with fields like
// name, age, and grade, and prints out the details.

package main

import "fmt"

type bioData struct {
	name  string
	age   int
	grade int
}

func struPoint(x *bioData) {

	fmt.Println("Student Details")
	fmt.Println(x.name)
	fmt.Println(x.age)
	fmt.Println(x.grade)

}

func main() {

	data := bioData{
		name:  "Philip",
		age:   7,
		grade: 4,
	}
	struPoint(&data)
}
