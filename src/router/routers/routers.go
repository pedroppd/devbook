package routers

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Router represent all routers from API
type Router struct {
	URI                    string
	Method                 string
	Func                   func(http.ResponseWriter, *http.Request)
	RequiredAuthentication bool
}

// SetUpRouter configure all routes within router
func SetUpRouter(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)
	for _, route := range routes {
		if route.RequiredAuthentication {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Func))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
		}
	}
	return r
}
