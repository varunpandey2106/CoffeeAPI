package helpers

import (
	"log"
	"os"
	"net/http"
	"encoding/json"
)

type Envelope map[string] interface{}

type Message struct {
	InfoLog*log.Logger
	Errorog *log.Logger
}


var infoLog=log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var errorLog=log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576
	r.Body=http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
    err := dec.Decode(data)
	if err != nil {
        return err
    }



}

