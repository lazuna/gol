package main

import (
	"fmt"
	"time"
)

func main() {
	// Reference time:
	// Mon Jan 2 04:04:04 MST 2006

	today := time.Now()
	fmt.Printf("It is currently %v\n", today.Format("Monday, Jan 2 2006"))
	startDate := time.Date(2020, 07, 04, 9, 00, 00, 0, time.UTC)
	elapsed := time.Since(startDate)

	fmt.Printf("%v has passed since %v\n", elapsed, startDate.Format("Monday, Jan 2 2006"))
	fmt.Printf("Hours: %v Minutes: %v Seconds: %v\n", elapsed.Hours(), elapsed.Minutes(), elapsed.Seconds())
	fmt.Printf("Hours: %.0f Minutes: %.0f Seconds: %.0f\n", elapsed.Hours(), elapsed.Minutes(), elapsed.Seconds())
}