package slack

import (
	"log"
	"os"
)

// ErrorLog
var ErrorLog *log.Logger

// InfoLog
var InfoLog *log.Logger

// EventLog
var EventLog *log.Logger

// Log
var Log *log.Logger

func init() {
	ErrorLog = log.New(os.Stderr, "[ERROR]\t", log.LstdFlags|log.Lshortfile)
	InfoLog = log.New(os.Stdout, "[INFO]\t", log.LstdFlags)
	EventLog = log.New(os.Stdout, "[EVENT]\t", log.LstdFlags)
	Log = log.New(os.Stdout, "[LOG]\t", log.LstdFlags)
}

// LogErrorsToFile redirects the ErrorLog output to the given file name
func LogErrorsToFile(filename string) {
	logToFile(ErrorLog, filename)
}

// LogEventsToFile redirects the EventLog output to the given file name
func LogEventsToFile(filename string) {
	logToFile(EventLog, filename)
}

// LogInfoToFile redirects the InfoLog output to the given file name
func LogInfoToFile(filename string) {
	logToFile(InfoLog, filename)
}

func logToFile(logger *log.Logger, filename string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatalf("error opening file: %v", err)
	}

	logger.SetOutput(f)
}
