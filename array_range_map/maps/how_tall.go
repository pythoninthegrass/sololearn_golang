package main
import (
	"fmt"
)

func main() {
    team := map[string]float32{
        "P1": 1.98,
        "P2": 2.05,
        "P3": 1.89,
        "P4": 2.0,
        "P5": 2.11}

	// get average height of all players
	var sum float32 = 0.0
	var count int = 0
	for _, v := range team {
		sum += v
		count++
	}
	fmt.Println(sum / float32(count))

	// * BONUS 🙌
	// f-string to print average height up to 2 decimal places (precision)
	fmt.Printf("Average height is %0.2f\n", sum / float32(count))

	// get the tallest player
	var player string
	var tallest float32 = 0.0
	for p, v := range team {
		if v > tallest {
			tallest = v
			player = p
		}
	}
	// print player map key with tallest height value
	fmt.Println("Tallest player is", player + " with height", tallest)
}
