package main

import (
	"cat-mail/src/config"
	"cat-mail/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()

	fmt.Println("Launching API...")
	fmt.Println("Creating routes...")
	r := router.CreateRoutes()
	fmt.Println("Routes created!")

	fmt.Printf("Listening on port: %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))

	//db := connection.Load()
	//defer db.Close()

}
