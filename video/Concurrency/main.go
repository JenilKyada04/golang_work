package main

import "fmt"

// go doSomething()

func febonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	fmt.Println("Concurrency in Go")

	c := make(chan int)
	go febonacci(10, c)
	for i := range c {
		fmt.Println(i)
	}
}
