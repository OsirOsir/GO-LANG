package main

import "fmt"

func main() {
	emplyeeData := newData()

	emplyeeData.updatebioFir("Philip")
	emplyeeData.updatebioseco("Osir")
	emplyeeData.updatebioAge(25)
	emplyeeData.updatebioGender("Male")
	emplyeeData.updatebioSalary(39000)

	fmt.Println(emplyeeData.format())
}
