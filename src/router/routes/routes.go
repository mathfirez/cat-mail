package routes

import (
	"cat-mail/src/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// API routes. All system routes are built using this struct.
type Route struct {
	Uri          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

// Place all routes in the router.
func Configure(r *mux.Router) *mux.Router {
	routes := msgRoutes

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}

	return r
}

// Defines all routes.
var msgRoutes = []Route{
	{
		Uri:          "/message",
		Method:       http.MethodPost,
		Function:     controllers.ParseMessage,
		AuthRequired: false,
	},
}
