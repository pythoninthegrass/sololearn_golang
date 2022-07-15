package main

import "fmt"

func square(x int) int {
	return x * x
}

func main() {
    var x int
	// fmt.Println("Enter a number: ")
    fmt.Scanln(&x)

	res := square(x)
   	fmt.Println(res)
}
