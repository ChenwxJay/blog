package ueditor

import (
	"net/http"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	action := r.URL.Query()["action"][0]
	if r.Method == "GET" {
		if action == "config" {
			config(w, r)
		}
		if action == "listimage" {
			listImage(w,r)
		}
	} else if r.Method == "POST" {
		if action == "uploadimage" {
			uploadImage(w, r)
		}
	}
}

