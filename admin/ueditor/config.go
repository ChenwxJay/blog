package ueditor

import (
	"bytes"
	"os"
	"net/http"
)

func config(w http.ResponseWriter, r *http.Request) {
	var file, err = os.Open("config.json")
	var configJson  []byte
	if err != nil {
		configJson = []byte("{}")
	} else {
		defer file.Close()
		buf := bytes.NewBuffer(nil)
		buf.ReadFrom(file)
		configJson = buf.Bytes()
	}
	w.Write(configJson)
}