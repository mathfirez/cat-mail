package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var ConnStr string
var Port int

// Loads env variables and connects to the DB. Must always be followed by a defer db.Close() when called.
func Load() (db *sql.DB) {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 8000
	}

	ConnStr = "postgres://postgres:postgres@localhost/socialNetwork?sslmode=disable"

	db, err = sql.Open("postgres", ConnStr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	//defer db.Close()
	return db
}
