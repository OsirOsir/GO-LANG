// 8. **Array Reversal:**
//    - Declare an array named `letters` containing 5 characters.
//    - Write a function to reverse the elements of the `letters` array in place and print the reversed array.

package main

import "fmt"

func reverseArrays(letters []rune) {

	for i, j := 0, len(letters)-1; i < j; i, j = i+1, j-1 {

		letters[i], letters[j] = letters[j], letters[i]
	}
}

func main() {
	letters := []rune{'P', 'h', 'i', 'l', 'i', 'O', 's', 'i', 'r'}

	reverseArrays(letters)
	fmt.Println(string(letters))
}
