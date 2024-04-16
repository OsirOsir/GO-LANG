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

		b.updatebioFir(firstName)
		b.updatebioseco(secondName)
		fmt.Println(firstName, secondName)
		promptOptions(b)

	case "2":
		age, _ := getInput("Age: ", reader)

		a, err := strconv.Atoi(age)
		if err != nil {
			fmt.Println("Age must be a number")
			promptOptions(b)
		}
		b.updatebioAge(a)
		fmt.Println(a)
		promptOptions(b)
	case "3":
		gender, _ := getInput("Gender: ", reader)
		fmt.Println(gender)
		b.updatebioGender(gender)
		promptOptions(b)
	case "4":
		salary, _ := getInput("Salary: ", reader)
		s, err := strconv.Atoi(salary)
		if err != nil {
			fmt.Println("Salary must be a number")
			promptOptions(b)
		}
		b.updatebioSalary(s)
		fmt.Println(s)
		promptOptions(b)
	case "5":
		b.save()
		fmt.Println("You saved the File")
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
