package main

import "fmt"

func main() {
	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			fmt.Println(i, " - ", j)
		}
	}
}
