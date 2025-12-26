package main

import "fmt"

// maps means key-value pair data structure
// golang struct means like object in other languages.

func main() {

	salaries := map[string]int{
		"Alice":  50000,
		"Bob":    56000,
		"Charli": 40000,
	}

	fmt.Println("This is salaray : ", salaries)

	fmt.Println("Alice salaray : ", salaries["Alice"])

	salaries["Bob"] = 800000
	fmt.Println("updated salary of Bob", salaries)

	if value, ok := salaries["Eve"]; ok {
		fmt.Println("Eve s salary ", value)
	} else {
		fmt.Println("Eve is not found in Map")
	}
	

	delete(salaries, "charlie")
	fmt.Println("Map after deleting charlie : ", salaries)
	fmt.Println("Total employees:", len(salaries))

	fmt.Printf("\n")

	fmt.Println("Interating over the map")
	for name, salary := range salaries {
		fmt.Printf("%s salary is %d \n" , name , salary)
	}

	fmt.Printf("\n")

	employeeID := make(map[string]int)
	employeeID["Frank"] = 101
	fmt.Println("Empty map created with make:", employeeID)

}
