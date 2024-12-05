package utils

import (
	"io"
	"log"
	"os"
	"tx-parser/config"
)

var (
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
	logFile     *os.File // File to hold the log file reference
)

func init() {
	var err error
	fileName := config.AppConfig.LogFile
	logFile, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}

	// Define loggers with prefixes and flags
	InfoLogger = log.New(io.MultiWriter(logFile, os.Stdout), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLogger = log.New(logFile, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// CloseLogFile safely closes the log file
func CloseLogFile() {
	if logFile != nil {
		logFile.Close()
	}
}
