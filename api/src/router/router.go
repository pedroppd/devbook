package router

import (
	"api/src/router/routers"

	"github.com/gorilla/mux"
)

//Gerar will return the router already configured
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return routers.SetUpRouter(r)
}
