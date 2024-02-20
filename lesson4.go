// Printing and Formatting  strings

package main

import "fmt"

func main() {

	age := 35
	name := "Philip"

	//Print

	fmt.Print("Hello \n")

	//Println
	fmt.Println("Hello Ninjas!")
	fmt.Println("my age is:", age, "and my name is:", name)

	//Printf (formating strings), %_ = format specifier

	fmt.Printf("my age is : %v and my name is %v: \n", age, name)
	fmt.Printf("my age is : %q and my name is %q: \n", age, name)
	fmt.Printf("my age is of type %T \n", age)
	fmt.Printf("my age is of type %T \n", age)
	fmt.Printf("I scroed %f \n", 34.6)
	fmt.Printf("I scroed %0.1f \n", 34.6)

	//Sprintf (saves formated strings )
	var str = fmt.Sprintf("my age is : %v and my name is %v: \n", age, name)

	fmt.Print("the stored data is ", str)

}
