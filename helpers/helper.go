package helpers

import (
	"log"
	"os"
)

type Envelope map[string] interface{}

type Message struct {
	InfoLog*log.Logger
	Errorog *log.Logger
}


var infoLog=log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var errorLog=log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

