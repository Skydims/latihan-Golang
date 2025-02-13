package main

import "fmt"

func main() {
	var x [58]string

	for i := 24; i <= 123; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)
	fmt.Println(x[42])
}
