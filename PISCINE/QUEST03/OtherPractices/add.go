package main

import "fmt"

func AddNum(a *int) {
	*a = *a + 1
}

func main() {
	a := 2
	AddNum(&a)
	fmt.Println(a)
}
