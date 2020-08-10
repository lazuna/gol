package main

import (
	"fmt"
)

type point struct {
	x, y int
}

type Person struct {
	firstName string
	lastname string
	age int
}

func main() {
	p := point{ 2, 3 }
	fmt.Printf("%v\n", p)
	fmt.Printf("%b\n", p) // binary
	fmt.Printf("%c\n", 33) // ASCII of number
	fmt.Printf("%x\n", 456) // to prune x values

	newPerson := Person{ "Raghav", "Dinesh", 20 }
	fmt.Printf("%v\n", newPerson)
	fmt.Printf("%T\n", newPerson)
}
