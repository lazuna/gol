package main

import (
	"fmt"
	"time"
)

const (
	formatISO = "2006-01-03"
	formatIN = "2 January, 2006"
	formatUS = "January 2, 2006"
)

func main() {
	t := time.Now()
	// Date has to be in this format to format it: Mon Jan 2 15:04 MST 2006
	fmt.Print(t, "\n")
	fmt.Print(t.Year(), "\n")
	fmt.Print(t.Month(), "\n")
	fmt.Print(t.Day(), "\n")

	Year := t.Year()
	Month := t.Month()
	Day := t.Day()

	fmt.Printf("Today is (MM/DD/YYYY) %d/%d/%d\n", Month, Day, Year)

	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Format(time.UnixDate))
	fmt.Println(t.Format(time.Kitchen))

	fmt.Println(t.Format("Monday, January 2 in the year 2006"))

	// Custom dates
	sDate := time.Date(2018, 07, 04, 9, 00, 00, 00, time.UTC)
	fmt.Println(t.Format("Monday, January 2 2006 at 15:04:05 2006"))

	fmt.Println(sDate.Format(formatISO))
	fmt.Println(sDate.Format(formatIN))
	fmt.Println(sDate.Format(formatUS))

}