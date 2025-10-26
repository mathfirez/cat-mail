package processor

import (
	"cat-mail/src/connection"
	"cat-mail/src/models"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
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

	userName = strings.ToLower(userName)

	var clientMessage models.ClientMessage

	fmt.Println(userName)
	// id will not be returned to the user. Will be used to update fields in the DB.
	row := db.QueryRow("SELECT id, author, posted_on, message from messages where receiver = $1 and printed = 'F' order by posted_on limit 1", userName)

	err := row.Scan(&clientMessage.Id, &clientMessage.Author, &clientMessage.PostedOn, &clientMessage.Content)

	if err != nil {
		if err == sql.ErrNoRows {
			procStatus := http.StatusNoContent
			return models.ClientMessage{}, procStatus
		} else {
			log.Println(err)
			procStatus := http.StatusBadRequest //400

			return models.ClientMessage{}, procStatus
		}

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

func GetClientIP(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	if ip == "::1" { //TODO: REMOVE -- TESTING ONLY
		ip = "localhost"
	}

	return ip
}

func RequestsCache(clientIP string) bool {
	fmt.Println("New request from:", clientIP)
	var tooManyRequests bool

	maxRequests, _ := strconv.Atoi(os.Getenv("API_PORT"))

	db := connection.Load()
	defer db.Close()

	var requestsCount int

	row := db.QueryRow("SELECT count(logged_on) from requests where ip = $1", clientIP)

	row.Scan(requestsCount) // will always return something

	if requestsCount >= maxRequests {
		tooManyRequests = true
		fmt.Println("Too many requests for", clientIP, ". Total: ", requestsCount)
		return tooManyRequests
	}

	_, err := db.Query("INSERT INTO requests (ip, logged_on) VALUES ($1, CURRENT_TIMESTAMP)", clientIP)

	if err != nil {
		log.Println("Unable to log request: ", err)
	} else {
		fmt.Println("Logged request for", clientIP)
	}

	return tooManyRequests
}
