package main

import (
	"fmt"
	"os"
	"strconv"
)

// To run this program, build the program instead of running .go file.
func main() {
	// To get all the arguments
	// args := os.Args
	// fmt.Println(args)

	// To start reading from argument 1
	args := os.Args[1:]

	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Usage: args <Food Cost> <Tip Percentage>")
		fmt.Println("Example: args 25 5")
	} else {
		if len(args) != 2 {
			fmt.Println("Enter min 2 arguments! & type /help for example.")
		} else {
			foodTotal, _ := strconv.ParseFloat(args[0], 32)
			tipAmount, _ := strconv.ParseFloat(args[1], 32)
			fmt.Printf("Total expenses: %.2f", expenseCalc(float32(foodTotal), float32(tipAmount)))
		}
	}
}

func expenseCalc(foodTotal float32, tipAmount float32) float32 {
	totalCost := foodTotal + (foodTotal * (tipAmount / 100))
	return totalCost
}
