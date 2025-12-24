package main

import "fmt"

// bool ,
// string ,
// int int8 int16 int32 int64 ,
// uint uint8 uint16 uint32 uint64,
// rune ,
// float32 ,float64 ,
// complex64 , complex128


// Assignment operator

func main() {
	a := "jenil"

	const fn = "a"
	const ln = "b"
	const fullname = fn + ln
	fmt.Println(fullname)

	// fmt.Println(a)
	fmt.Printf("The type is : %T\n ", a)

	mess := 10
	mess1 := 12

	if mess > mess1 {
		fmt.Println("Message sent ")
	} else {
		fmt.Println("Message is not sent")
	}

}
