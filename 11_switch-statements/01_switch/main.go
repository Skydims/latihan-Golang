package main

import "fmt"

func main() {
	switch "Mhi" {
	case "Dimas":
		fmt.Println("Wassup Dimas")
	case "Samid":
		fmt.Println("Wassup Samid")
	case "Mimin":
		fmt.Println("Wassup Mimin")
	default:
		fmt.Println("Have you no friends?")
	}
}

/*
  no default fallthrough
  fallthrough is optional
  -- you can specify fallthrough by explicitly stating it
  -- break isn't needed like in other languages
*/
