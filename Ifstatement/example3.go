package main

import "fmt"

func main() {
	if age := 51; age%2 == 0 { //checks if number is even
		fmt.Println(age, "is even")
	} else {
		fmt.Println(age, "is odd")
	}
}

// In the above program age is initialized in the if statement.
//One thing to be noted is that age is available only for access from inside the if and else.
// i.e. the scope of age is limited to the if else blocks.
// If we try to access num from outside the if or else, the compiler will complain.
// This syntax often comes in handy when we declare a variable just for the purpose of if else construct.
// Using this syntax in such cases ensures that the scope of the variable is only within the if else statement.
