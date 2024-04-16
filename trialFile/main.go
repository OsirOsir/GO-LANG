package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createData() bioData {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Add Personal Details: ")
	// data, _ := reader.ReadString('\n')
	// data = strings.TrimSpace(data)

	data, _ := getInput("Add Personal Details: ", reader)

	b := newData()

	fmt.Println("Personal Data added", data)
	return b
}

func promptOptions(b bioData) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose which detail to add: 1 Names:, 2 Age:, 3 Gender:, 4 Salary:, 5 SaveData: ", reader)
	// fmt.Println("You chose:", opt)

	switch opt {
	case "1":
		firstName, _ := getInput("First Name: ", reader)
		secondName, _ := getInput("Second Name: ", reader)
		fmt.Println(firstName, secondName)
		promptOptions(b)

	case "2":
		age, _ := getInput("Age: ", reader)

		a, err := strconv.ParseFloat(age, 64)
		if err != nil {
			fmt.Println("Age must be a number")
			promptOptions(b)
		}
		fmt.Println(a)
		promptOptions(b)
	case "3":
		gender, _ := getInput("Gender: ", reader)
		fmt.Println(gender)
		promptOptions(b)
	case "4":
		salary, _ := getInput("Salary: ", reader)
		s, err := strconv.ParseFloat(salary, 64)
		if err != nil {
			fmt.Println("Salary must be a number")
			promptOptions(b)
		}
		fmt.Println(s)
		promptOptions(b)
	case "5":
		fmt.Println("Data saved")
	default:
		fmt.Println("Invalid Choice")
		promptOptions(b)
	}
}

func main() {

	employeeData := createData()
	promptOptions(employeeData)
	// emplyeeData := newData()

	// emplyeeData.updatebioFir("Philip")
	// emplyeeData.updatebioseco("Osir")
	// emplyeeData.updatebioAge(25)
	// emplyeeData.updatebioGender("Male")
	// emplyeeData.updatebioSalary(39000)

	// fmt.Println(employeeData)
}
