package processor

import (
	"cat-mail/src/connection"
	"cat-mail/src/models"
	"fmt"
	"log"
	"net/http"
)

// Adds a message to the queue (db) and returns the corresponding http status code depending on the status of processing.
// TODO - Address different error codes to return.
func AddToQueue(message models.Message) int {

	db := connection.Load()
	defer db.Close()

	_, err := db.Query("INSERT INTO MESSAGES (author, receiver, message, posted_on, printed) VALUES ($1, $2, $3, CURRENT_TIMESTAMP, $4)", message.Author, message.Receiver, message.Content, "F")

	if err != nil {
		log.Println(err)
		procStatus := http.StatusBadRequest //400

		return procStatus
	}

	fmt.Println("Posted!")
	procStatus := http.StatusCreated //201

	return procStatus
}

// TODO Add comment
func GetMessageFromUser(userName string) (models.ClientMessage, int) {
	db := connection.Load()
	defer db.Close()

	var clientMessage models.ClientMessage

	fmt.Println(userName)
	// id will not be returned to the user. Will be used to update fields in the DB.
	row := db.QueryRow("SELECT id, author, posted_on, message from messages where receiver = $1 and printed = 'F' order by posted_on limit 1", userName)

	err := row.Scan(&clientMessage.Id, &clientMessage.Author, &clientMessage.PostedOn, &clientMessage.Content)

	if err != nil {
		log.Println(err)
		procStatus := http.StatusBadRequest //400

		return models.ClientMessage{}, procStatus
	}

	procStatus := http.StatusOK

	return clientMessage, procStatus
}

// TODO - Improve and Review
func LogMessageSent(id string) {
	db := connection.Load()
	defer db.Close()

	_, err := db.Query("UPDATE MESSAGES SET printed = 'T', printed_on = CURRENT_TIMESTAMP WHERE id = $1", id)

	if err != nil {
		log.Println(err)
	}
}
