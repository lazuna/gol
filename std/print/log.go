package main

import (
	"runtime"
	"fmt"
)

func main() {
	// This does not insert a line break!
	fmt.Print("Current version of Go is: ", runtime.Version())
	// This does insert a line break!
	fmt.Print("\n")

	// This does insert a line break!
	fmt.Println("Current version of Go is: ", runtime.Version())

	// This does not insert a line break!
	// %v is value
	fmt.Printf("Current version of Go is: %v\n", runtime.Version())
	fmt.Printf("Current version of Go is: %v and OS is: %v\n", runtime.Version(), runtime.GOOS)
}