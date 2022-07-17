package main
import "fmt"

type Cart struct {
    prices []float32
}

func (x Cart) show() {
    var sum float32 = 0.0

    //calculate the sum of all prices in the Cart
	for _, v := range x.prices {
		sum += v
	}
    fmt.Println(sum)
}

func main() {
	c := Cart{}
	var n int
	fmt.Scanln(&n)

	// take n inputs and add them to the Cart
	for i := 0; i < n; i++ {
		var price float32
		fmt.Scanln(&price)
		c.prices = append(c.prices, price)
	}

  	//call the show() method of the Cart
	c.show()
}
