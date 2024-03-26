package structFile

type Address struct {
	Country  string
	Province string
}

type Details struct {
	Name   string
	Age    int
	Gender string
	Height int
	Weight float64
	Address
}
