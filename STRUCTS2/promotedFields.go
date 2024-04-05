package main

import "fmt"

type SalaryData struct {
	bankName  string
	accNumber int
	wage      int
}

type BioData struct {
	name   string
	age    int
	status string
	SalaryData
}

func main() {
	EmlInfo := BioData{
		name:   "Philip Osir Otieno",
		status: "Married",
		SalaryData: SalaryData{
			bankName:  "Kenya Commercial Bank",
			accNumber: 47844458564456,
			wage:      67000,
		},
	}

	fmt.Println("Name :", EmlInfo.name)
	fmt.Println("Age :", EmlInfo.age)
	fmt.Println("Status :", EmlInfo.status)
	fmt.Println("Bank Name :", EmlInfo.bankName)
	fmt.Println("Account Number :", EmlInfo.accNumber)
	fmt.Printf("Salary: $%d\n", EmlInfo.wage)

}
