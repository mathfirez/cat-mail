package processor

import (
	"cat-mail/src/models"
	"fmt"
	"time"
)

// Adds a message to the queue (db).
// TODO
func AddToQueue(message models.Message) {
	fmt.Println(message)
}

// TODO: Deprecate here, will be moved to the client.
// Queries the database for messages and sends to the printer queue based on the interval defined in the .env file.
// Sends a single message per run to avoid bloating.
func Scheduler(interval int) {
	seconds := time.Duration(interval) * time.Second
	for {
		time.Sleep(seconds)
		//Querying new messages

		// If no messages, sleep again
		// Else, sends to printer, update message fields on DB based (see models.Message)
	}

}
