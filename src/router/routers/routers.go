package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Router represent all routers from API
type Router struct {
	URI                    string
	Method                 string
	Func                   func(http.ResponseWriter, *http.Request)
	RequiredAuthentication bool
}

//SetUpRouter configure all routes within router
func SetUpRouter(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}
	return r
}
