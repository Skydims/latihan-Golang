package main

import "fmt"

func main() {
	x := 24
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "aduh "
		fmt.Println(y)
	}
	// fmt.Println(y) // outside scope of y
}
