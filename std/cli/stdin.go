package main

import (
	"fmt"
	"os"
	"bufio"
)
func main() {
	// body
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Who are you?")

	// '\n' Expects the user to press Enter key after inserting value
	text, _ := reader.ReadString('\n')
	fmt.Printf("Hey %v", text)
}