package logger

import (
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
	Info.SetOutput(os.Stdout)
	Error.SetOutput(os.Stderr)
	Debug.SetOutput(os.Stdout)
	log.SetOutput(Debug.Writer())
	log.SetPrefix("DEBUG ")
	log.SetFlags(flags)
}
