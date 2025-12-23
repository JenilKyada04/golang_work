package main

import (
	"fmt"
	"maps"
)

// Maps used to store collections of unorded key and values pairs.

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 10

	fmt.Println("map", m)

	v1 := m["k1"]
	fmt.Println("v1", v1)

	v2 := m["k3"]
	fmt.Println("v2", v2)

	fmt.Println("len : ", len(m))

	delete(m, "k2")
	fmt.Println(m)
	clear(m)
	fmt.Println("map:", m)

	n := map[string]int{"foo": 4, "ram": 19}
	fmt.Println("Map : ", n)

	n1 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n1) {
		fmt.Println("n == n2")
	}

}
