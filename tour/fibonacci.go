package main

import "fmt"

func febonacii() func() int {
	a, b := 0, 1
	return func() int {
		current := a
		a, b = b, a+b
		return current
	}
}

func main() {
	f := febonacii()
	for i := 0; i < 5; i++ {
		fmt.Println(f())
	}

}
