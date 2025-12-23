package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "JENIL"
	fmt.Println(len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}
	fmt.Println()

	fmt.Println("Runn count ", utf8.RuneCountInString(s))

	for idx, runsValue := range s {
		fmt.Printf("%#U start at %d \n", runsValue, idx)
	}

	// fmt.Println()

	for i, w := 0, 0; i < len(s); i += w {
		runsValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U start at %d \n ", runsValue, i)
		w = width
	}

}
