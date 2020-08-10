package main

import (
	"fmt"
)

func main() {
	// %4.2f represents 4 digits before the decimal and 2 after it.
	// Alignment left %7.2f (positive values)
	// Alignment right %-7.2f (negative values)
	// It works same way on strings too. or you can mix string and floats values as well.
	fmt.Printf("Right Justified: |%7.2f|%7.2f|%7.2f|\n", 23.3774, 577.45, 1234.56)
	fmt.Printf("Left  Justified: |%-7.2f|%-7.2f|%-7.2f|\n", 98.999,12.3456, 12.01)
}