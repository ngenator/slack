package slack

import (
	"io"
	"log"
	"os"
)

type Log struct {
	*log.Logger
	enabled bool
}

func (l *Log) ToFile(filename string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		l.Logger.Fatalf("error opening file: %v", err)
	}
	l.Logger.SetOutput(f)
}

func (l *Log) Output(calldepth int, s string) error {
	if l.enabled {
		return l.Logger.Output(calldepth, s)
	}
	return nil
}

func (l *Log) Disable() {
	l.enabled = false
}

func (l *Log) Enable() {
	l.enabled = true
}

func NewLog(out io.Writer, prefix string, flag int) *Log {
	return &Log{
		log.New(out, prefix, flag),
		true,
	}
}

var (
	ErrorLog *Log
	InfoLog  *Log
	DebugLog *Log
	EventLog *Log
)

func init() {
	ErrorLog = NewLog(os.Stderr, "[ERROR]\t", log.LstdFlags|log.Lshortfile)
	InfoLog = NewLog(os.Stdout, "[INFO]\t", log.LstdFlags)
	DebugLog = NewLog(os.Stdout, "[DEBUG]\t", log.LstdFlags)
	EventLog = NewLog(os.Stdout, "[EVENT]\t", log.LstdFlags)
}
