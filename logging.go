package slack

import (
	"log"
	"os"
)

var (
	ErrorLog *log.Logger
	EventLog *log.Logger
	Log      *log.Logger
)

func init() {
	ErrorLog = log.New(os.Stderr, "[ERROR]\t", log.LstdFlags|log.Lshortfile)
	EventLog = log.New(os.Stdout, "[EVENT]\t", log.LstdFlags)
	Log = log.New(os.Stdout, "[INFO]\t", log.LstdFlags)
}

func LogErrorsToFile(filename string) {
	logToFile(ErrorLog, filename)
}

func LogEventsToFile(filename string) {
	logToFile(EventLog, filename)
}

func LogInfoToFile(filename string) {
	logToFile(Log, filename)
}

func logToFile(logger *log.Logger, filename string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatalf("error opening file: %v", err)
	}

	logger.Println("Output redirecting to ", filename)
	logger.SetOutput(f)
	logger.Println("Output changed!")
}
