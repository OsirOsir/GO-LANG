// 3. **Array Iteration:**
//    - Declare an array named `ages` containing the ages of 5 people.
//    - Use a loop to iterate over the `ages` array and print each age.

// package main

// import "fmt"

// func main() {
// 	ages := [...]int{20, 35, 21, 45, 67}

// 	for _, v := range ages{
// 		fmt.Println(v)
// 	}
// }

package main

import "fmt"

func main() {
	ages := [...]int{34, 25, 76, 45, 32}
	for i := 0; i <= 4; i++ {
		fmt.Println(ages[i])
	}
}
