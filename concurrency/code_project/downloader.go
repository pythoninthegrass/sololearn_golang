package main
import "fmt"

// take the size of the file as an argument and
// count the sum of all integers from 0 to that number (inclusive)
func download(url int, ch chan int) {
	var sum int
	for i := 0; i <= url; i++ {
		sum += i
	}
	ch <- sum
}

func main() {
	// create a channel to receive the sum of all integers from 0 to 100
    ch1 := make(chan int)
    ch2 := make(chan int)
    ch3 := make(chan int)

	// takes three file sizes as inputs (e.g., 10, 100, 42)
    var s1, s2, s3 int
	fmt.Printf("Enter a file size (e.g., 10): \n")
    fmt.Scanln(&s1)
    fmt.Scanln(&s2)
    fmt.Scanln(&s3)

	// calls the download() function as Goroutines for each file
    go download(s1, ch1)
    go download(s2, ch2)
    go download(s3, ch3)

    // output the sum of all results (e.g., 6008)
	fmt.Println(<-ch1 + <- ch2 + <- ch3)
}
