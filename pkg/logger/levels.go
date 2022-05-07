package logger

import (
	"io"
	"log"
	"os"
)

const flags = log.Ldate | log.Ltime | log.Lshortfile

var (
	Info  = log.New(os.Stdout, "INFO  ", flags)
	Debug = log.New(os.Stdout, "DEBUG ", flags)
	Error = log.New(os.Stderr, "ERROR ", flags)
)

func init() {
	logFile, err := os.OpenFile("fahro.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		Error.Fatalln("Could not initialize log file:", err)
	}
	// write to stdout and fahro.log
	output := io.MultiWriter(os.Stdout, logFile)
	Info.SetOutput(output)
	Error.SetOutput(output)

	logFile, err = os.OpenFile("fahro.debug.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		Error.Fatalln("Could not initialize debug log file:", err)
	}
	Debug.SetOutput(io.MultiWriter(os.Stdout, logFile))

	log.SetOutput(Debug.Writer())
	log.SetPrefix("DEBUG ")
	log.SetFlags(flags)
}
