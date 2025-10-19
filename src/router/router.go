package router

import (
	"cat-mail/src/router/routes"

	"github.com/gorilla/mux"
)

func CreateRoutes() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
