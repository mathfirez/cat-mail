package main

import (
	"cat-mail/src/connection"
	"cat-mail/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	connection.Load()

	fmt.Println("Launching API...")
	fmt.Println("Creating routes...")
	r := router.CreateRoutes()
	fmt.Println("Routes created!")

	fmt.Printf("Listening on port: %d", connection.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", connection.Port), r))

	//db := connection.Load()
	//defer db.Close()

}
