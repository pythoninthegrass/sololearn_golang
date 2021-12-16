package main

import (
	"fmt"
//	"strings"
)

func main() {
//	fmt.Println(strings.Repeat("GO\n", 3))
    x := 6
    y := x-1
    x += y
    x %= y
    fmt.Println(x)
}
