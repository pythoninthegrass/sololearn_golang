package main

import "fmt"

func repeat(word string, number int) {
	for i := 0; i < number; i++ {
		fmt.Println(word)
	}
}

func main() {
    var w string
	// fmt.Print("Enter a word: ")
	fmt.Scan(&w)

	var x int
	// fmt.Print("Enter a number: ")
	fmt.Scan(&x)

	repeat(w, x)
}
