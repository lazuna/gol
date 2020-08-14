package main

import (
	"fmt"
	"time"
)

func main() {

	firstStamp := time.Now()
	fmt.Printf("It is at the moment %v \n", firstStamp.Format("15:04:05"))
	// time.Sleep(2000000000) // time in nano-second
	// time.Sleep(time.Second) // runs for one second
	time.Sleep(2 * time.Second) // runs for two second
	secondStamp := time.Now()
	fmt.Printf("Time stamp now %v\n", secondStamp.Format("15:04:05"))

}
