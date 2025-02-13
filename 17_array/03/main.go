package main

import "fmt"

func main() {
	var x [241]int

	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 0; i < 241; i++ {
		x[i] = i
	}
	for i, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 45 {
			break
		}
	}
}
