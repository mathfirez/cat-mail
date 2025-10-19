package controllers

import (
	"cat-mail/src/models"
	"cat-mail/src/processor"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// ProcessMessage
func ProcessMessage(w http.ResponseWriter, r *http.Request) {
	bodyReq, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var message models.Message
	if err = json.Unmarshal(bodyReq, &message); err != nil {
		log.Fatal(err)
	}

	processor.AddToQueue(message)

	w.Write([]byte("TO-DO"))
}
