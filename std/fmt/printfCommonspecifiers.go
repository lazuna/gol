package main

import (
	"fmt"
)

func main() {
	var age int = 10
	var name string = "Raghav Dinesh"
	/*
		%v - formats the value in a default formats
		%s - formats string values
		%q - formats string values with quotes
		%d - formats decimal integers
		%g - formats the floating-point numbers
		%b - formats base 22 numbers
		%o - formats base 88 numbers
		%t - formats true or false (boolean) values
	*/

	fmt.Printf("My name is %s and I am %d years old\n", name, age)
}

