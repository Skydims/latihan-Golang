package main

import (
	"errors"
	"log"
)

func main() {
	_, err := Sqrt(-12)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// implementation
	return 42, nil
}
