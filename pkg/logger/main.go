package logger

import (
	"log"
	"os"
)

func NewLogger(typeLogger string) *log.Logger {
	f := os.Stderr
	if typeLogger == "ERROR" || typeLogger == "WARNING" {
		return log.New(f, typeLogger+": ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		return log.New(f, typeLogger+": ", log.Ldate|log.Ltime)
	}
}
