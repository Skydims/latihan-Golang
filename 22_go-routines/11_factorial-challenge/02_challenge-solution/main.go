package main

import (
	"fmt"
)

func main() {
	c := factorial(4)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for j := n; j > 0; i-- {
			total *= j
		}
		out <- total
		close(out)
	}()
	return out
}
