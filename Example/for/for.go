package main

import "fmt"

func main() {

	i := 0
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}

	for n := 0; n < 6; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
