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
		RequiredAuthentication: true,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodGet,
		Func:                   controllers.FindUser,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodPut,
		Func:                   controllers.UpdateUser,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodDelete,
		Func:                   controllers.DeleteUser,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/users/{id}/follow",
		Method:                 http.MethodPost,
		Func:                   controllers.FollowUser,
		RequiredAuthentication: true,
	},
}
