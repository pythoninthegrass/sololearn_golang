package main
import (
	"fmt"
	"strings"
)

func main() {
	// w: 3
	// d: 1
	// l: 0
    results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}

	// take the last match result as input and append it to the results array. After that, calculate and output the total points the team gained from the results.
	var lastMatchResult string
	// fmt.Printf("You winning, son? (w/d/l)\n")
	fmt.Scanln(&lastMatchResult)
	results = append(results, lastMatchResult)

	// TODO: refactor as case statement
	var totalPoints int
	for _, result := range results {
		// lowercase the result
		result = strings.ToLower(result)
		if result == "w" {
			totalPoints += 3
		} else if result == "d" {
			totalPoints += 1
		} else {
			totalPoints += 0
		}
	}
	fmt.Println(totalPoints)
}
