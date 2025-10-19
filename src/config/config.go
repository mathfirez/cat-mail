package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var DBConnStr string
var Port int

// Initializes environment variables.
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 8000
	}

	// DB params.
	DBConnStr = fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PW"),
		os.Getenv("API_PORT"),
		os.Getenv("DB_NAME"),
	)
}
