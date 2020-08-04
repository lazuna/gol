package main

import (
	"fmt"
	"flag"
)

func main() {
	// It takes x86 by default.
	archPtr := flag.String("arch", "x86", "CPU Type")
	flag.Parse()

	switch *archPtr {
		case "x86":
			fmt.Println("Running the program in 32 bit mode.")
		case "AMD64":
			fmt.Println("Running the program in 64 bit mode.")
		case "IA64":
			fmt.Println("Eh! What's IA64?")
	}

	fmt.Println("Process complete.")
}