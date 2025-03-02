package routers

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Router{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Func:                   controllers.CreateUser,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Func:                   controllers.FindUsers,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodGet,
		Func:                   controllers.FindUser,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodPut,
		Func:                   controllers.UpdateUser,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodDelete,
		Func:                   controllers.DeleteUser,
		RequiredAuthentication: false,
	},
}
