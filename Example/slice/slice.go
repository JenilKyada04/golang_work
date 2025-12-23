package main

import (
	"fmt"
	// "slices"
)

func main() {

	var s []string
	fmt.Println("uninit", s, s == nil, len(s) == 0)

	fmt.Println("DONE")

	// c := make([]string, len(s))
	// copy(c, s)
	// fmt.Println("cpy:", c)

	// l := s[2:5]
	// fmt.Println("sl1:", l)

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
