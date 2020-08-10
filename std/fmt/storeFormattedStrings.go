package main

import (
	"fmt"
)

func main() {
	output := fmt.Sprintf("|%-7s|%-7s|%-7s|", "foo", "bar", "go")
	print(output)
}