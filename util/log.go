package util

import (
	"errors"
	"fmt"
	"log"
)

// LogFatalErr - Print a log message and return er the error
func LogFatalErr(message string) error {
	log.Fatal(message)
	return errors.New(message)
}

// LogFatalErrf - Print a log message and return an error wit this value
func LogFatalErrf(message string, args ...interface{}) error {
	log.Fatalf(message, args)
	return fmt.Errorf(message, args)
}

// LogPrintf - Print a formatted message
func LogPrintf(tag string, args ...interface{}) {
	log.Printf(fmt.Sprintf("%s ", tag), args)
}
