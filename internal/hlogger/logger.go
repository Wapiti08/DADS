package hlogger

import (
	"log"
	"os"
	"sync"
)

type hLogger struct {
	*log.Logger
	filename string
}

var hlogger *hLogger
var once sync.Once

// create a singleton instance of logger
func GetInstance() *hLogger {
	// create file for once, append mode for context
	once.Do(
		func() {
			hlogger = createLogger()
		}
	)
}

// create a logger instance
func createLogger(fname string) *hLogger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &hLogger{
		filename: fname,
		Logger: log.New(file, "DADS ", log.Lshortfile),
	}
}