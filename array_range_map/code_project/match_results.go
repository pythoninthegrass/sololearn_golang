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

	// Take the last match result as input and append it to the results array.
	// After that, calculate and output the total points the team gained from the results.
	var lastMatchResult string
	// fmt.Printf("You winning, son? (w/d/l)\n")
	fmt.Scanln(&lastMatchResult)
	results = append(results, lastMatchResult)

	var totalPoints int
	for _, result := range results {
		// lowercase the result
		result = strings.ToLower(result)
		switch result {
		case "w":
			totalPoints += 3
		case "d":
			totalPoints += 1
		default:
			totalPoints += 0
		}
	}
	fmt.Println(totalPoints)
}
