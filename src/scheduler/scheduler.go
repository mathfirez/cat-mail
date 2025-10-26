package scheduler

import (
	"cat-mail/src/connection"
	"fmt"
	"time"
)

func ClearRequests(interval int, minSecondsToClear int) {
	for {
		time.Sleep(time.Duration(interval) * time.Second)
		fmt.Println("Scheduler: Running...")

		db := connection.Load()

		_, err := db.Exec("DELETE FROM requests WHERE EXTRACT(EPOCH FROM current_timestamp - logged_on) > $1 ", minSecondsToClear)

		if err != nil {
			fmt.Println("Scheduler: Could not clear requests...")
		} else {
			fmt.Println("Scheduler: Requests cleared!")
		}

		db.Close()
	}
}
