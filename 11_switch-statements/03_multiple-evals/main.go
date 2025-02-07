package main

import "fmt"

func main() {
	switch "Dimas" {
	case "Tim", "Dimas":
		fmt.Println("Wassup Tim, or, err, dimas")
	case "Marcus", "Medhi":
		fmt.Println("Both of your names start with M")
	case "Julian", "Sushant":
		fmt.Println("Wassup Julian / Sushant")
	}
}
