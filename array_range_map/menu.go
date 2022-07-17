package main
import (
	"fmt"
)

func main() {
	// create array with n strings
  	menu := [6]string{"Water", "Burger", "Cake", "Soup", "Soda", "Fries"}

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
