package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 2
	}()
	fmt.Println(<-c)
}
