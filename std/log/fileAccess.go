package main

import (
	"log"
	"os"
)

type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
	FATAL
)

func main() {
	// file, err := os.Open("log.txt") // normal way to open files
	/*
	// Moved this block of code into a function 
	file, err := os.OpenFile("./assets/log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("This is log message!")
	*/

	writeLog(INFO, "This is an information message!")
	writeLog(WARNING, "This is a warning message!")
	writeLog(ERROR, "This is an error message!")
	writeLog(FATAL, "This is an crash message!")
	writeLog(INFO, "This message will never see the light!")
}

func writeLog(messagetype messageType, message string) {

	file, err := os.OpenFile("./assets/log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println(messagetype)
	switch messagetype {
		case INFO:
			logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
			logger.Println(message)
		case WARNING:
			logger := log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
			logger.Println(message)
		case ERROR:
			logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
			logger.Println(message)
		case FATAL:
			logger := log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
			logger.Println(message)
	}

}