package main
import "fmt"

// TODO: refactor to be variadic (has to be last argument in function)
// Take a number N as input, followed by N strings, which represent city names.
// Store the city names in a slice.
func route(nums int, cities []string) {
	// output the cities in the slice as a single string (e.g., Boston->Chicago->Washington->)
	var cityString string
	for i := 0; i < nums; i++ {
		cityString += cities[i] + "->"
	}
	fmt.Println(cityString)
}


func main() {
    var n int
    fmt.Scanln(&n)

    // take n strings as input and append them to the slice
	var cities []string
	for i := 0; i < n; i++ {
		var city string
		fmt.Scanln(&city)
		cities = append(cities, city)
	}

	route(n, cities)
}
