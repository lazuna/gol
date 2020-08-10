package main

import (
	"fmt"
)

func main() {
	// Printf formats according to a format specifier and writes standard output.
	// It returns the number of bytes written and any write error encountered.
	func Printf(format string, a ...interface{}) (n int, err error) {
		return Fprintf(os.Stdout, format, a...)
	}
}