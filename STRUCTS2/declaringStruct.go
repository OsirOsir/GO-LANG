package main

type employee struct {
	firstName string
	lastName  string
	age       int
	weight    float64
}

// other ways of  declaring a struct

type employee struct {
	firstName, lastName string
	age                 int
	weight              int
}
