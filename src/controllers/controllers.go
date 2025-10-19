package controllers

import "net/http"

// ParseMessage
func ParseMessage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TO-DO"))
}
