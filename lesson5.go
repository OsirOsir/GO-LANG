package main

import "fmt"

func main() {
	var age [5]int = [5]int{24, 25, 38, 45, 89}
	var age2 = [4]int{34, 54, 65, 67}

	names := [2]string{"Alice", "Philip"}

	alph := [5]rune{'A', 'B', 'C', 'D', 'E'}

	fmt.Println(age, age2, len(age), len(age2))
	fmt.Println(names, len(names))
	fmt.Printf("%c\n", alph)
	fmt.Println(len(alph))

}
