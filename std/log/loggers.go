package main

import (
	"log"
	"os"
)

var (
	WarningLogger	*log.Logger
	InfoLogger		*log.Logger
	ErroLogger		*log.Logger
	FatalLogger		*log.Logger
)

func init() {

	file, err := os.OpenFile("./assets/log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(file, "INFO (Additional Info): ", log.LUTC|log.Lmicroseconds|log.Llongfile) // Bit different logger.
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErroLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	FatalLogger = log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	InfoLogger.Println("This is info")
	WarningLogger.Println("Oh no, a warning!")
	ErroLogger.Println("Error! Please Stop")
	FatalLogger.Println("Program Exited")
}