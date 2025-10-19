package controllers

import (
	"net/http"
	"strings"
)

// HandleMessage
func HandleMessage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TO-DO"))
}

func ProcessMessage(raw_message string) (string){
	return strings.Trim(raw_message, " ")
}
