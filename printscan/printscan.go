package main

import (
	"fmt"
)

func main() {
	var yourNumber int
	fmt.Println("Please enter a number")
	fmt.Scan(&yourNumber)
	fmt.Printf("Your number was %v\n", yourNumber)
}
