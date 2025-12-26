package main

import "fmt"

// cap means capacity
// cap function return last element index + 1.
// ...int means any number of int arguments

// func printstrings(string ...string) {
// 	for i := 0; i < len(string); i++ {
// 		fmt.Println(string[i])
// 	}
// }

func main() {
	numbers := []int{1, 2, 3}
	fmt.Printf("numbers: %v, len: %d, cap: %d\n", numbers, len(numbers), cap(numbers))

	primes := [6]int{2, 3, 5, 7, 11, 13}
	slice := primes[1:4]
	fmt.Printf("slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

	s := make([]string, 3, 5)
	fmt.Printf("s: %v, len: %d, cap: %d\n", s, len(s), cap(s))

	a := []string{"a", "b", "c"}
	b := append(a, "d", "e")
	fmt.Printf("a: %v, len: %d, cap: %d\n", a, len(a), cap(a))
	fmt.Printf("b: %v, len: %d, cap: %d\n", b, len(b), cap(b))

	// names := []string{"a", "b"}
	// fmt.Println(names...)
	
}
