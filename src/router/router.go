package router

import (
	"cat-mail/src/router/routes"

	"github.com/gorilla/mux" // One of the popular routing frameworks
)

func CreateRoutes() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
