package main

import "fmt"

func main() {
	switch "Marcus" {
	case "DIMAS":
		fmt.Println("Wassup DIMAS")
	case "SAMID":
		fmt.Println("Wassup SAMID")
	case "MIMIN":
		fmt.Println("Wassup MIMIN")
		fallthrough
	case "SARAH":
		fmt.Println("Wassup SARAH")
		fallthrough
	case "JULIAN":
		fmt.Println("Wassup Julian")
	case "SARI":
		fmt.Println("Wassup SARI")
	}
}
