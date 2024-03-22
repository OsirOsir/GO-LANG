package main

import "fmt"

func main() {
	messagesFromDoris := []string{
		"You doing anything later??",
		"Did you get my last Message?",
		"Dont leave me hanging...",
		"Please respond I'm lonely",
	}

	numMessages := float64(len(messagesFromDoris))
	costPerMessage := .02

	totals := costPerMessage * numMessages

	fmt.Printf("The tota for all the Messages %.2f\n", totals)
}
