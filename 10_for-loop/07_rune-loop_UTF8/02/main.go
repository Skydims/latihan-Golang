package main

import "fmt"

func main() {
	for i := 12; i <= 124; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}
