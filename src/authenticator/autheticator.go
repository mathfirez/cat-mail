package authenticator

import (
	"cat-mail/src/connection"
	"log"
	"net/http"
)

// TODO - I dont know how to do this will just query the db for now lol
func Authenticate(headerToken string) (string, int) {
	db := connection.Load()
	defer db.Close()

	var userName string
	row := db.QueryRow("select name from clients where token = $1", headerToken)

	err := row.Scan(userName)

	if err != nil {
		log.Println(err)
		procStatus := http.StatusBadRequest //400

		return "", procStatus
	}

	procStatus := http.StatusOK

	return userName, procStatus
}
