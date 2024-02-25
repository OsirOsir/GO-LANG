// 4. **Array Update:**
//    - Declare an array named `scores` containing 3 integers representing test scores.
//    - Update the second element of the `scores` array to `85`.
//    - Print the updated `scores` array.


ackage main

import "fmt"

func main() {
	scores := [...]int{67, 89, 56}

	updatedScores := scores
	updatedScores[1] = 85

	fmt.Println(updatedScores)
}

// package main

// import "fmt"

// func main() {
// 	scores := [...]int{67, 89, 56}

// 	updatedScores := scores
// 	updatedScores[1] = 85

// 	for _, v := range updatedScores {
// 		fmt.Println(v)
// 	}
// }
