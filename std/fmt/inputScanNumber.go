package main

import (
	"fmt"
)

func main() {
	var fNumber int
	var sNumber int
	fmt.Println("What two numbers would you like to add?")
	fmt.Scanf("%d %d", &fNumber, &sNumber)
	fmt.Printf("Total is: %d", fNumber + sNumber)
}