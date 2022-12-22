package router

import (
	"github.com/Michael19s/go_first_api/controller"
	"github.com/labstack/echo/v4"
)

func registerUserRouter(e *echo.Echo) {
	e.GET("/users", controller.UserGetAll)
	e.GET("/users/:id", controller.UserGetById)
	e.POST("/users", controller.UserCreate)
	e.PUT("/users", controller.UserUpdate)
	e.DELETE("/users/:id", controller.UserDelete)
}
