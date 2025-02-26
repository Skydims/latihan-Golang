package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		if i >= 24 {
			break
		}
		i++
	}
}
