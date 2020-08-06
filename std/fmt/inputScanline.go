package main

import (
	"fmt"
)

func main () {
	fmt.Print("What is your name?")
	var fname string
	var lname string
	fmt.Scanln(&fname, &lname)
	fmt.Printf("Hello %s %s nice to meet you!\n", fname, lname)
}