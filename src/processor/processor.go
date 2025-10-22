package processor

import (
	"cat-mail/src/connection"
	"cat-mail/src/models"
	"fmt"
	"log"
)

// Adds a message to the queue (db).
func AddToQueue(message models.Message) {

	db := connection.Load()
	defer db.Close()

	qry := "INSERT INTO MESSAGES (author, message, posted_on, printed) VALUES ($1, $2, CURRENT_TIMESTAMP, $3)"

	_, err := db.Exec(qry, message.Author, message.Content, "F")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Posted!")

}
