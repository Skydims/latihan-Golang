package main

import "fmt"

func main() {
	var x [243]string

	fmt.Println(len(x))
	fmt.Println(x[0])
	for i := 0; i < 243; i++ {
		x[i] = string(i)
	}
	for _, v := range x {
		fmt.Printf("%v - %T - %v\n", v, v, []byte(v))
	}
}
