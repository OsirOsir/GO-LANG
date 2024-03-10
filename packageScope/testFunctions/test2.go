package main

func sum(b ...int) int {
	var total int
	for _, v := range b {
		total += v
	}
	return total
}
