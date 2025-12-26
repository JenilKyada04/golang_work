package main

import (
	"errors"
	"fmt"
	"log"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	var num1, num2 int

	fmt.Print("Enter two integers (separated by a space): ")

	_, err := fmt.Scanln(&num1, &num2)
	if err != nil {
		log.Fatal("Invalid input:", err)
	}

	result, err := Divide(num1, num2)
	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Println("Result:", result)
}
