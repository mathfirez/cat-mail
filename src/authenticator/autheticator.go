package authenticator

import "cat-mail/src/connection"

//TODO
func Authenticate() bool {
	db := connection.Load()
	defer db.Close()

	return true
}
