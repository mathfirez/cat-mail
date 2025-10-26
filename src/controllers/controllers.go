package controllers

import (
	"cat-mail/src/authenticator"
	"cat-mail/src/models"
	"cat-mail/src/processor"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Process incoming messages from the web page.
func ProcessMessage(w http.ResponseWriter, r *http.Request) {
	bodyReq, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var message models.Message
	if err = json.Unmarshal(bodyReq, &message); err != nil {
		log.Fatal(err)
	}

	procStatus := processor.AddToQueue(message)

	w.WriteHeader(procStatus)
}

// Authenticates the token and queries for the pending messages for the user related to the token. Replies with the message and http status code.
// TODO - add comments for readability. Review.
func GetMessage(w http.ResponseWriter, r *http.Request) {
	clientIP := processor.GetClientIP(r) //TODO Error handling

	tooManyRequests := processor.RequestsCache(clientIP)

	if tooManyRequests {
		w.WriteHeader(http.StatusTooManyRequests)
		return
	}

	headerToken := r.Header.Get("Authorization")

	if headerToken == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	userName, statusCode := authenticator.Authenticate(headerToken) // TODO

	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
		return
	}

	clientMessage, procStatus := processor.GetMessageFromUser(userName)

	if procStatus != http.StatusOK {
		w.WriteHeader(procStatus)
		return
	}

	clientMessageJson, err := json.Marshal(clientMessage)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest) // TODO - Not sure if the right thing to send
		return
	}
	// TODO Update Message fields in the DB.
	id := clientMessage.Id
	processor.LogMessageSent(id)

	// Clear messageID
	clientMessage.Id = ""

	w.WriteHeader(procStatus)
	w.Write(clientMessageJson)
}
