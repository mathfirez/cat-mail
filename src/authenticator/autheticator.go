package authenticator

import (
	"cat-mail/src/connection"
	"net/http"
)

// TODO
func Authenticate(headerToken string) (string, int) {
	db := connection.Load()
	defer db.Close()

	return "mathfirez", http.StatusOK
}
