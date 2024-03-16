package main

import "fmt"

type bioStatus struct {
	age    int
	status string
}

func main() {
	agent := map[string]bioStatus{
		"Naomi":  {25, "single"},
		"Philip": {30, "Maried"},
		"Joshua": {42, "Divoced"},
		"Marion": {32, "Maried"},
	}
	newAgent := "Alice"
	_, k := agent[newAgent]
	if k == true {
		fmt.Printf("The name %s is available\n", newAgent)
	} else {
		fmt.Println(newAgent, "not available")
	}
	for k, v := range agent {
		fmt.Printf("The agent %s, is %d years old, and is %s\n", k, v.age, v.status)
	}
}
