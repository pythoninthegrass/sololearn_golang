package main

// ! SoloLearn doesn't like 'golang.org' imports: `go.mod file not found in current directory or any parent`
import (
	"fmt"
	"strings"
	// "golang.org/x/text/cases"
	// "golang.org/x/text/language"
)

func num_to_word(num int) string {
	switch num {
	case 0:
		return "zero"
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	default:
		return "Not a number between 0 and 10"
	}
}

func main() {
	// fmt.Println("Enter 3 numbers between 0 and 10: ")
	var num1, num2, num3 int

	// newline between each input
	fmt.Scan(&num1, &num2, &num3)

	// convert nums to words
	word1 := num_to_word(num1)
	word2 := num_to_word(num2)
	word3 := num_to_word(num3)

	// * FANCY IMPORTS (i.e., correct library)
	// print out the words as title case
	// fmt.Printf("%s\n%s\n%s\n",
	// 	cases.Title(language.Und).String(
	// 		word1),
	// 	cases.Title(language.Und).String(
	// 		word2),
	// 	cases.Title(language.Und).String(
	// 		word3))

	// ! DEPRECATED METHOD
	// print out the words as title case
	fmt.Printf("%s\n%s\n%s\n",
		strings.Title(word1),
		strings.Title(word2),
		strings.Title(word3))
}
