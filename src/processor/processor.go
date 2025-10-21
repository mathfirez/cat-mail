package processor

import (
	"cat-mail/src/models"
	"fmt"
)

// Adds a message to the queue (db).
// TODO
func AddToQueue(message models.Message) {
	fmt.Println(message)
}
