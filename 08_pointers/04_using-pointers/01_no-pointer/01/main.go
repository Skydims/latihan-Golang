package main

import "fmt"

func zero(z int) {
	z = 0
}

func main() {
	x := 4
	zero(x)
	fmt.Println(x) // x is still 4
}
