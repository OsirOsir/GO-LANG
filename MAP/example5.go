package main

import "fmt"

func main() {
	areaValue := make(map[string]float64)
	areaValue["Circle"] = 34.9
	areaValue["Rectanguler"] = 56.1
	areaValue["Triangle"] = 76.2

	for k, v := range areaValue {
		fmt.Println("Area of ", k, "=", v)

	}

	areaValue["Squire"] = 90.8
	areaValue["Oval"] = 45
	fmt.Println(areaValue)

	newArea := "Trapeziam"

	value, ok := areaValue[newArea]
	if ok == true {
		fmt.Println("The area of ", newArea, "is", value)
	} else {
		fmt.Println(newArea, "Not found")
	}

}
