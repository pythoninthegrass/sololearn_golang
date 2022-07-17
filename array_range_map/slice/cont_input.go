package main
import "fmt"

func main() {
	// take a number, indicating the number of inputs,
	var n int
	fmt.Scanln(&n)

	// create a slice of length n
	x := make([]int, n)

	// appends the inputs to the slice
	for i := 0; i < n; i++ {
		fmt.Scanln(&x[i])
	}

	// print inputs
	fmt.Println(x)
}
