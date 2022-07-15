package main

import (
	"fmt"
)

// go uses structs instead of classes
type Employee struct {
	name string
	age  int
	position string
	salary float32
}

func main() {
    e1 := Employee{"James", 42, "Manager", 0}

    var x float32
	// fmt.Print("Enter a salary for this poor bastard: ")
    fmt.Scanln(&x)
    e1.salary = x

    fmt.Println("========")
    fmt.Println(e1.name + "(" + e1.position + ")") // [sic: no space between name and (]
    fmt.Println(e1.age)
    fmt.Println(e1.salary)
    fmt.Println("========")
}
