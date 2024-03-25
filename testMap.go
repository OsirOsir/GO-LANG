// package main

// import "fmt"

// func main() {
// 	employeSalaries := map[string]int{
// 		"Philip": 34000,
// 		"Erick":  50000,
// 		"Elvis":  90000,
// 	}

// 	employeSalaries["Fred"] = 24000

// 	for k, v := range employeSalaries {
// 		fmt.Println(k, v)
// 	}
// 	newPerson := "Elvis"
// 	// value, ok := map[key]
// 	value, ok := employeSalaries[newPerson]

// 	if ok == true {
// 		fmt.Printf("Slary of %s is %v\n", newPerson, value)
// 	} else {
// 		fmt.Println("Not found")
// 	}

// 	// delete(map, key)

// 	delete(employeSalaries, "Elvis")
// 	fmt.Println(employeSalaries)
// }

package main

import "fmt"

type employee struct {
	age     int
	salary  int
	country string
}

func main() {
	employeeDetails := map[string]employee{
		"Philip": {34, 34000, "Kenya"},
		"Joshua": {28, 46999, "Uganda"},
		"Albert": {23, 25777, "Cameroon"},
	}
	for k, v := range employeeDetails {
		fmt.Printf("Name = %s\n Age = %d\n Salary = %v\n Country = %s\n", k, v.age, v.salary, v.country)
	}
}
