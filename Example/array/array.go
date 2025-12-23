package main

import "fmt"

func main() {
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d : ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{2, 3, 4},
	}

	fmt.Println("2d : ", twoD)

	b := [...]int{100, 3: 100, 200}
	fmt.Println("Idx : ", b)

}
