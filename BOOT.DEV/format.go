package main

import "fmt"

func main() {
	fname := "Dalinar"
	lname := "Kholin"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."

	userLog := fmt.Sprintf("Name: %s %s, Age: %d, Rate: %.1f, isSubscribed: %t, Message: %v", fname, lname, age, messageRate, isSubscribed, message)

	// Don't touch above this line

	fmt.Println(userLog)
}
