package main

import "fmt"

func main() {
	colors := []string{"blue", "green", "red"}
	for index, value := range colors {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	for _, value := range colors {
		fmt.Println("Value:", value)
	}

	a := 10
	for a < 20 {
		a++
		fmt.Println(a)
	}

	// modual
	// divide Two type of operation

	// logical 
	// Assignment
	
}
