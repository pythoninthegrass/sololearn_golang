package main

import "fmt"

// timer base types
type Timer struct {
	id string
	value int
}

// tick function in place with pointer
func (t *Timer) tick(val int) {
	t.value += val
	fmt.Println(t.value)
}

// timer is named "timer1", empty initial value, increment by 1
func main() {
    var x int
    fmt.Scanln(&x)

    t := Timer{"timer1", 0}

    for i:=0;i<x;i++ {
        t.tick(1)
    }
}
