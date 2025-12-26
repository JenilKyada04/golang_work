package main

import "fmt"

func main() {
    num := 42
    fmt.Println("Initial value of num:", num)
    fmt.Println("Memory address of num:", &num) 

    var ptr *int = &num

    fmt.Println("\nValue of pointer (the address):", ptr)
    fmt.Println("Value at address (dereferenced):", *ptr)

    *ptr = 100 

    fmt.Println("\nValue after modification through pointer:", *ptr)
    fmt.Println("Original value of num is also changed:", num)
}
