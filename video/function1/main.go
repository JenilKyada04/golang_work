
package main

import "fmt"

// func concat(s1 string, s2 string) string {
// 	return s1 + s2
// }

func increment(a, b int) int {
	return a + b
}
func getname() (string, string) {
	return "JENIL", "KYADA"
}

func main() {

	// fmt.Println(concat("Hello ", "World"))

	firstname, lastname := getname()

	fmt.Println(firstname , lastname)

	fmt.Println(increment(5, 10))

}

