package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "/help" {
		fmt.Println("Usage: fileReader <input file>")
	} else {
		// Help text.
		fmt.Println("How do you wish to format your text?")
		fmt.Println("1: ALL CAPS")
		fmt.Println("2: Title Case")
		fmt.Println("3: lower case")

		// Scan user choice
		var option int
		_, err := fmt.Scanf("%d", &option)

		// Opens the file.
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			switch option {
				case 1:
					fmt.Println(strings.ToUpper(scanner.Text()))
				case 2:
					fmt.Println(strings.Title(scanner.Text()))
				case 3:
					fmt.Println(strings.ToLower(scanner.Text()))
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}
}