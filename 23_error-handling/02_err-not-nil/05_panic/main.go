package main

import (
	"os"
)

func main() {
	_, err := os.Open("noo-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		//		log.Fatalln(err)
		panic(err)
	}
}
