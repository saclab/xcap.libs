package logger

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	c "github.com/TreyBastian/colourize"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		c.Colourize("[TRACE] ", c.Magenta),
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		c.Colourize("[INFO] ", c.Cyan),
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		c.Colourize("[WARNING] ", c.Yellow),
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		c.Colourize("[FATAL] ", c.Red, c.Blinkslow),
		log.Ldate|log.Ltime|log.Lshortfile)
}

func InitD() {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
}
