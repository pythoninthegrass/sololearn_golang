package main

import (
	"fmt"
)

// scale function in place with pointer
func scale(num *float32, factor float32) {
	*num *= factor
}

func main() {
	// num := float32(4)
	// factor := float32(1.5)

	// input num and factor
	var num float32
	// fmt.Print("Enter a number: ")
    fmt.Scan(&num)

	var factor float32
	// fmt.Print("Enter a factor (e.g., 1.5): ")
    fmt.Scan(&factor)

	// scale num in place with pointer
	scale(&num, factor)

	// print result
	// fmt.Printf("%s %f\n", "The scaled number is", num)
	fmt.Println(num)
}
