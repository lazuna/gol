package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("What's your name?")

	// %s - To read normal string.
	// inputs, _ := fmt.Scanf("%s", &name)

	// %q - To read quoted string.
	inputs, _ := fmt.Scanf("%q", &name)

	switch inputs {
		case 0:
			fmt.Printf("You must enter a name!\n")
		case 1:
			fmt.Printf("Hello %s! Nice to meet you\n", name)
	}

}
