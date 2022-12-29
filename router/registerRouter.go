package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandlerFunc(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "WELCOME!!")
}

// Register here all the routes to be used in the API
func RegisterRouter(e *echo.Echo) {
	e.GET("/", HandlerFunc)
	registerUserRouter(e)
}
