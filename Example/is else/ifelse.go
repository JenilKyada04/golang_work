package main

import "fmt"

func main() {

	var id int

	fmt.Print("Enter the value for check even and odd number : ")
	fmt.Scanf("%d", &id)

	if id%2 == 0 {
		fmt.Println(id, "is even")
	} else {
		fmt.Println(id, "is odd")
	}

}
