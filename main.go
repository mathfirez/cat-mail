package main

import (
	"cat-mail/src/config"
	"cat-mail/src/processor"
	"cat-mail/src/router"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	config.Load()

	fmt.Println("Launching API...")
	fmt.Println("Creating routes...")
	r := router.CreateRoutes()
	fmt.Println("Routes created!")

	// Gets Scheduler interval.
	interval, err := strconv.Atoi(os.Getenv("INTERVAL"))

	if err != nil {
		log.Fatal(err)
	}

	// Starts scheduler service.
	go processor.Scheduler(interval)

	fmt.Printf("Listening on port: %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))

	//db := connection.Load()
	//defer db.Close()

}
