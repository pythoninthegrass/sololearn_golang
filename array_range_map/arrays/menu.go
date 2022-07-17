package main
import (
	"fmt"
)

func main() {
	// * Better OG option (i.e., Occam's Razor)
	// create array with n strings
  	// menu := [6]string{"Water", "Burger", "Cake", "Soup", "Soda", "Fries"}

	// ! Feature-complete option
	// 33.1 slices lesson
	// make creates array of type slice with a zero placeholder for the length
	menu := make([]string, 0)
	// append adds an element to the end of the slice
	menu = append(menu, "Water", "Burger", "Cake", "Soup", "Soda", "Fries")
	// fmt.Println(menu[5])

	// assign user input to variable
	var n int
  	// fmt.Print("Enter a number between 0 and 5: ")
	fmt.Scanln(&n)

	// n between 0 and 5
	// if index is invalid, return "Invalid choice"
	if n < 0 || n > 5 {
		fmt.Println("Invalid choice")
	} else {
		fmt.Println(menu[n])
	}
}
