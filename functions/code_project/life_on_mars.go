package main

import (
	"fmt"
)

type Age struct {
	age int
	yr int
}

func (a *Age) mars_age(val int) {
	earth_yr := 365
	mars_yr := 687

	a.age = val * earth_yr
	a.yr = a.age / mars_yr
}

func main() {
    var age int
	// fmt.Println("Enter your age (years): ")
    fmt.Scanln(&age)

    mars := Age{0, 0}
	mars.mars_age(age)
	// fmt.Printf("Your age on Mars:\n%d\n", mars.yr)
	fmt.Println(mars.yr)
}
