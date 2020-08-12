package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	"runtime/trace"
)

func main() {

	/*
	** To read trace.out file you need to run a tool using following command
	** > go tool trace trace.out
	** Which opens in a new tab in your browser http://127.0.0.1:59912/
	*/
	f, err := os.Create("./assets/trace.out")
	if err != nil {
		log.Fatalf("Couldn't create a trace file! %v\n", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("Couldn't close the file! %v\n", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("Failed to start a trace: %d\n", err)
	}

	defer trace.Stop()

	AddRandomNumbers()
}

func AddRandomNumbers() {
	firstNumber := rand.Intn(100)
	secondNumber := rand.Intn(100)

	time.Sleep(2 * time.Second)

	var result = firstNumber + secondNumber
	fmt.Printf("Result of 2 numbers is %d\n", result)
}