package main

import (
	"fmt"
)

func main() {
	var fname string
	var lname string
	fmt.Println("What is your name?")
	// Because values in Scanf are space separeted!
	fmt.Scanf("%s %s", &fname, &lname)
	fmt.Printf("Hello %s %s! Nice to meet you!\n", fname, lname)
}