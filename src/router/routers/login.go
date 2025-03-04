package routers

import (
	"api/src/controllers"
	"net/http"
)

var loginRoute = Router{
	URI:                    "/login",
	Method:                 http.MethodPost,
	Func:                   controllers.Login,
	RequiredAuthentication: false,
}
