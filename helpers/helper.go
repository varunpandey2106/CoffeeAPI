package helpers

import "log"

type Envelope map[string] interface{}

type Message struct {
	InfoLog*log.Logger
	Errorog *log.Logger
}

