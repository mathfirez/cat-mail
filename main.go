package main

import (
	"cat-mail/src/connection"
	"cat-mail/src/router"
	"cat-mail/src/scheduler"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	connection.Load()

	interval, _ := strconv.Atoi(os.Getenv("SCHEDULER_INTERVAL"))
	minSecondsToClear, _ := strconv.Atoi(os.Getenv("SCHEDULER_MIN_MINUTES_FROM_LAST_REQUEST"))

	go scheduler.ClearRequests(interval, minSecondsToClear)

	fmt.Println("Launching API...")
	fmt.Println("Creating routes...")
	r := router.CreateRoutes()
	fmt.Println("Routes created!")

	fmt.Printf("Listening on port: %d\n", connection.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", connection.Port), r))

	//db := connection.Load()
	//defer db.Close()

}
