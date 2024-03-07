package main

func calculateAvarage(n []float64) float64 {
	if len(n) == 0 {
		return 0
	}
	var total float64
	for _, v := range n {
		total += v
	}
	var avg float64
	avg = total / float64(len(n))
	return avg
}
