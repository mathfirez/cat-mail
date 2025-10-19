package controllers

import "net/http"

// ProcessMessage
func ProcessMessage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TO-DO"))
}
